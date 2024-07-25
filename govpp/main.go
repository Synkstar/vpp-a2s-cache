package main

import (
	"bytes"
	"encoding/binary"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/rumblefrog/go-a2s"

	"github.com/joho/godotenv"

	vppa2s "govpp/binapi/a2s"

	"net"

	"go.fd.io/govpp"
	"go.fd.io/govpp/adapter/socketclient"
	"go.fd.io/govpp/core"
)

var (
	socketAddr = flag.String("sock", socketclient.DefaultSocketName, "Path to binary API socket file")
)

type Server struct {
	address string
	port    int
	strict  bool
}

func main() {
	flag.Parse()
	godotenv.Load()

restart:
	for {
		time.Sleep(time.Second)
		conn, connEv, err := govpp.AsyncConnect(*socketAddr, core.DefaultMaxReconnectAttempts, core.DefaultReconnectInterval)
		if err != nil {
			log.Printf("ERROR Connecting to api: %s\n", err)
			continue
		}
		//defer conn.Disconnect()

		e := <-connEv
		if e.State != core.Connected {
			log.Printf("Error: connecting to VPP failed %s\n", e.Error)
			continue
		}
		ch, err := conn.NewAPIChannel()
		if err != nil {
			log.Printf("Error: creating channel failed %s\n", err)
			continue
		}
		//defer ch.Close()

		hostsEnv := os.Getenv("HOSTS")
		if hostsEnv == "" {
			fmt.Println("HOSTS environment variable not set")
			return
		}

		updateIntervalEnv := os.Getenv("UPDATE_INTERVAL")
		updateInterval, err := strconv.Atoi(updateIntervalEnv)
		if err != nil {
			fmt.Printf("Update interval not set or invalid")
			return
		}

		hosts := strings.Split(hostsEnv, "),(")

		var servers []Server
		for _, host := range hosts {
			host = strings.Trim(host, "()")
			parts := strings.Split(host, ",")
			if len(parts) != 2 {
				fmt.Printf("Invalid host format: %s\n", host)
				continue
			}

			addressPort := strings.Trim(parts[0], `"`)
			strictModeStr := strings.Trim(parts[1], `"`)
			strictMode, err := strconv.ParseBool(strictModeStr)
			if err != nil {
				fmt.Printf("Invalid strict mode value: %s\n", strictModeStr)
				continue
			}

			addressParts := strings.Split(addressPort, ":")
			if len(addressParts) != 2 {
				fmt.Printf("Invalid address format: %s\n", addressPort)
				continue
			}

			port, err := strconv.Atoi(addressParts[1])
			if err != nil {
				fmt.Printf("Invalid port value: %s\n", addressParts[1])
				continue
			}

			server := Server{
				address: addressParts[0],
				port:    port,
				strict:  strictMode,
			}
			servers = append(servers, server)
		}

		for {
			for _, server := range servers {
				host := fmt.Sprintf("%s:%d", server.address, server.port)
				client, err := a2s.NewClient(host, a2s.SetMaxPacketSize(14000), a2s.TimeoutOption(time.Second*5))
				if err != nil {
					fmt.Printf("Setting up query client for %s failed\n", host)
					continue
				}
				info, err := client.QueryInfo()
				if err != nil {
					fmt.Printf("A2S_INFO query to %s failed\n", host)
					continue
				}
				players, err := client.QueryPlayer()
				if err != nil {
					fmt.Printf("A2S_PLAYER query to %s failed\n", host)
					continue
				}
				rules, err := client.QueryRules()
				if err != nil {
					fmt.Printf("A2S_RULES query to %s failed\n", host)
					continue
				}
				parts := strings.Split(host, ":")
				if len(parts) != 2 {
					fmt.Printf("expected format 'IP:PORT', got %s\n", host)
					continue
				}

				ip := net.ParseIP(parts[0]).To4()
				if ip == nil {
					fmt.Printf("Invalid ip address: %s\n", parts[0])
					continue
				}
				ipIntHostOrder := binary.LittleEndian.Uint32(ip)

				port, err := strconv.Atoi(parts[1])
				if err != nil {
					fmt.Printf("Invalid port %s\n", parts[1])
					continue
				}

				key := vppa2s.A2sKey{
					IP:   ipIntHostOrder,
					Port: uint16(port),
				}

				infoBuf := new(bytes.Buffer)
				binary.Write(infoBuf, binary.LittleEndian, &info.Data)
				infoBytes := infoBuf.Bytes()
				infoLength := len(infoBytes)

				playersBuf := new(bytes.Buffer)
				binary.Write(playersBuf, binary.LittleEndian, &players.Data)
				playersBytes := playersBuf.Bytes()
				playersLength := len(playersBytes)

				rulesBuf := new(bytes.Buffer)
				binary.Write(rulesBuf, binary.LittleEndian, &rules.Data)
				rulesBytes := rulesBuf.Bytes()
				rulesLength := len(rulesBytes)

				value := vppa2s.A2sData{InfoData: infoBytes, PlayerData: playersBytes, RulesData: rulesBytes, InfoLength: uint16(infoLength), PlayerLength: uint16(playersLength), RulesLength: uint16(rulesLength), Strict: server.strict}

				kv := vppa2s.ClibBihashKvA2s{
					Key:   key,
					Value: value,
				}

				request := vppa2s.A2sSetData{
					IsAdd: true,
					Kv:    kv,
				}
				reply := &vppa2s.A2sSetDataReply{}

				if err := ch.SendRequest(&request).ReceiveReply(reply); err != nil {
					fmt.Printf("Error sending request: %v\n", err)
					goto restart
				}

				fmt.Printf("Successfully updated %s\n", host)
			}
			time.Sleep(time.Duration(updateInterval) * time.Second)
		}
	}
}
