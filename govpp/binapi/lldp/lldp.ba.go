// Code generated by GoVPP's binapi-generator. DO NOT EDIT.
// versions:
//  binapi-generator: v0.10.0
//  VPP:              24.10-rc0~84-gd2d41bc34
// source: plugins/lldp.api.json

// Package lldp contains generated bindings for API file lldp.api.
//
// Contents:
// -  2 enums
// -  7 messages
package lldp

import (
	"govpp/binapi/interface_types"
	"govpp/binapi/ip_types"
	"strconv"

	api "go.fd.io/govpp/api"
	codec "go.fd.io/govpp/codec"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the GoVPP api package it is being compiled against.
// A compilation error at this line likely means your copy of the
// GoVPP api package needs to be updated.
const _ = api.GoVppAPIPackageIsVersion2

const (
	APIFile    = "lldp"
	APIVersion = "2.0.0"
	VersionCrc = 0xc26a6a68
)

// ChassisIDSubtype defines enum 'chassis_id_subtype'.
type ChassisIDSubtype uint32

const (
	CHASSIS_ID_SUBTYPE_RESERVED     ChassisIDSubtype = 0
	CHASSIS_ID_SUBTYPE_CHASSIS_COMP ChassisIDSubtype = 1
	CHASSIS_ID_SUBTYPE_INTF_ALIAS   ChassisIDSubtype = 2
	CHASSIS_ID_SUBTYPE_PORT_COMP    ChassisIDSubtype = 3
	CHASSIS_ID_SUBTYPE_MAC_ADDR     ChassisIDSubtype = 4
	CHASSIS_ID_SUBTYPE_NET_ADDR     ChassisIDSubtype = 5
	CHASSIS_ID_SUBTYPE_INTF_NAME    ChassisIDSubtype = 6
	CHASSIS_ID_SUBTYPE_LOCAL        ChassisIDSubtype = 7
)

var (
	ChassisIDSubtype_name = map[uint32]string{
		0: "CHASSIS_ID_SUBTYPE_RESERVED",
		1: "CHASSIS_ID_SUBTYPE_CHASSIS_COMP",
		2: "CHASSIS_ID_SUBTYPE_INTF_ALIAS",
		3: "CHASSIS_ID_SUBTYPE_PORT_COMP",
		4: "CHASSIS_ID_SUBTYPE_MAC_ADDR",
		5: "CHASSIS_ID_SUBTYPE_NET_ADDR",
		6: "CHASSIS_ID_SUBTYPE_INTF_NAME",
		7: "CHASSIS_ID_SUBTYPE_LOCAL",
	}
	ChassisIDSubtype_value = map[string]uint32{
		"CHASSIS_ID_SUBTYPE_RESERVED":     0,
		"CHASSIS_ID_SUBTYPE_CHASSIS_COMP": 1,
		"CHASSIS_ID_SUBTYPE_INTF_ALIAS":   2,
		"CHASSIS_ID_SUBTYPE_PORT_COMP":    3,
		"CHASSIS_ID_SUBTYPE_MAC_ADDR":     4,
		"CHASSIS_ID_SUBTYPE_NET_ADDR":     5,
		"CHASSIS_ID_SUBTYPE_INTF_NAME":    6,
		"CHASSIS_ID_SUBTYPE_LOCAL":        7,
	}
)

func (x ChassisIDSubtype) String() string {
	s, ok := ChassisIDSubtype_name[uint32(x)]
	if ok {
		return s
	}
	return "ChassisIDSubtype(" + strconv.Itoa(int(x)) + ")"
}

// PortIDSubtype defines enum 'port_id_subtype'.
type PortIDSubtype uint32

const (
	PORT_ID_SUBTYPE_RESERVED         PortIDSubtype = 0
	PORT_ID_SUBTYPE_INTF_ALIAS       PortIDSubtype = 1
	PORT_ID_SUBTYPE_PORT_COMP        PortIDSubtype = 2
	PORT_ID_SUBTYPE_MAC_ADDR         PortIDSubtype = 3
	PORT_ID_SUBTYPE_NET_ADDR         PortIDSubtype = 4
	PORT_ID_SUBTYPE_INTF_NAME        PortIDSubtype = 5
	PORT_ID_SUBTYPE_AGENT_CIRCUIT_ID PortIDSubtype = 6
	PORT_ID_SUBTYPE_LOCAL            PortIDSubtype = 7
)

var (
	PortIDSubtype_name = map[uint32]string{
		0: "PORT_ID_SUBTYPE_RESERVED",
		1: "PORT_ID_SUBTYPE_INTF_ALIAS",
		2: "PORT_ID_SUBTYPE_PORT_COMP",
		3: "PORT_ID_SUBTYPE_MAC_ADDR",
		4: "PORT_ID_SUBTYPE_NET_ADDR",
		5: "PORT_ID_SUBTYPE_INTF_NAME",
		6: "PORT_ID_SUBTYPE_AGENT_CIRCUIT_ID",
		7: "PORT_ID_SUBTYPE_LOCAL",
	}
	PortIDSubtype_value = map[string]uint32{
		"PORT_ID_SUBTYPE_RESERVED":         0,
		"PORT_ID_SUBTYPE_INTF_ALIAS":       1,
		"PORT_ID_SUBTYPE_PORT_COMP":        2,
		"PORT_ID_SUBTYPE_MAC_ADDR":         3,
		"PORT_ID_SUBTYPE_NET_ADDR":         4,
		"PORT_ID_SUBTYPE_INTF_NAME":        5,
		"PORT_ID_SUBTYPE_AGENT_CIRCUIT_ID": 6,
		"PORT_ID_SUBTYPE_LOCAL":            7,
	}
)

func (x PortIDSubtype) String() string {
	s, ok := PortIDSubtype_name[uint32(x)]
	if ok {
		return s
	}
	return "PortIDSubtype(" + strconv.Itoa(int(x)) + ")"
}

// configure global parameter for LLDP
//     - system_name - VPP system name
//     - tx_hold - multiplier for tx_interval when setting time-to-live (TTL)
//                      value in the LLDP packets
//     - tx_interval - time interval, in seconds, between each LLDP frames
// LldpConfig defines message 'lldp_config'.
type LldpConfig struct {
	TxHold     uint32 `binapi:"u32,name=tx_hold" json:"tx_hold,omitempty"`
	TxInterval uint32 `binapi:"u32,name=tx_interval" json:"tx_interval,omitempty"`
	SystemName string `binapi:"string[],name=system_name" json:"system_name,omitempty"`
}

func (m *LldpConfig) Reset()               { *m = LldpConfig{} }
func (*LldpConfig) GetMessageName() string { return "lldp_config" }
func (*LldpConfig) GetCrcString() string   { return "c14445df" }
func (*LldpConfig) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *LldpConfig) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4                     // m.TxHold
	size += 4                     // m.TxInterval
	size += 4 + len(m.SystemName) // m.SystemName
	return size
}
func (m *LldpConfig) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeUint32(m.TxHold)
	buf.EncodeUint32(m.TxInterval)
	buf.EncodeString(m.SystemName, 0)
	return buf.Bytes(), nil
}
func (m *LldpConfig) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.TxHold = buf.DecodeUint32()
	m.TxInterval = buf.DecodeUint32()
	m.SystemName = buf.DecodeString(0)
	return nil
}

// LldpConfigReply defines message 'lldp_config_reply'.
type LldpConfigReply struct {
	Retval int32 `binapi:"i32,name=retval" json:"retval,omitempty"`
}

func (m *LldpConfigReply) Reset()               { *m = LldpConfigReply{} }
func (*LldpConfigReply) GetMessageName() string { return "lldp_config_reply" }
func (*LldpConfigReply) GetCrcString() string   { return "e8d4e804" }
func (*LldpConfigReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *LldpConfigReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Retval
	return size
}
func (m *LldpConfigReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	return buf.Bytes(), nil
}
func (m *LldpConfigReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	return nil
}

// Details about neighbor
//     - sw_if_index - interface where neighbor was discovered
//     - last_heard - last heard time
//     - last_sent - last sent time
//     - chassis_id - chassis id value
//     - chassis_id_len - length for chassis id
//     - port_id - port id value
//     - port_id_len - length for port id
//     - ttl - time to length for the neighbour
//     - port_id_subtype - subtype for port_id
//     - chassis_id_sybtype - sybtype for chassis_id
// LldpDetails defines message 'lldp_details'.
// InProgress: the message form may change in the future versions
type LldpDetails struct {
	SwIfIndex        interface_types.InterfaceIndex `binapi:"interface_index,name=sw_if_index" json:"sw_if_index,omitempty"`
	LastHeard        float64                        `binapi:"f64,name=last_heard" json:"last_heard,omitempty"`
	LastSent         float64                        `binapi:"f64,name=last_sent" json:"last_sent,omitempty"`
	ChassisID        []byte                         `binapi:"u8[64],name=chassis_id" json:"chassis_id,omitempty"`
	ChassisIDLen     uint8                          `binapi:"u8,name=chassis_id_len" json:"chassis_id_len,omitempty"`
	PortID           []byte                         `binapi:"u8[64],name=port_id" json:"port_id,omitempty"`
	PortIDLen        uint8                          `binapi:"u8,name=port_id_len" json:"port_id_len,omitempty"`
	TTL              uint16                         `binapi:"u16,name=ttl" json:"ttl,omitempty"`
	PortIDSubtype    PortIDSubtype                  `binapi:"port_id_subtype,name=port_id_subtype" json:"port_id_subtype,omitempty"`
	ChassisIDSubtype ChassisIDSubtype               `binapi:"chassis_id_subtype,name=chassis_id_subtype" json:"chassis_id_subtype,omitempty"`
}

func (m *LldpDetails) Reset()               { *m = LldpDetails{} }
func (*LldpDetails) GetMessageName() string { return "lldp_details" }
func (*LldpDetails) GetCrcString() string   { return "c2d226cd" }
func (*LldpDetails) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *LldpDetails) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4      // m.SwIfIndex
	size += 8      // m.LastHeard
	size += 8      // m.LastSent
	size += 1 * 64 // m.ChassisID
	size += 1      // m.ChassisIDLen
	size += 1 * 64 // m.PortID
	size += 1      // m.PortIDLen
	size += 2      // m.TTL
	size += 4      // m.PortIDSubtype
	size += 4      // m.ChassisIDSubtype
	return size
}
func (m *LldpDetails) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeUint32(uint32(m.SwIfIndex))
	buf.EncodeFloat64(m.LastHeard)
	buf.EncodeFloat64(m.LastSent)
	buf.EncodeBytes(m.ChassisID, 64)
	buf.EncodeUint8(m.ChassisIDLen)
	buf.EncodeBytes(m.PortID, 64)
	buf.EncodeUint8(m.PortIDLen)
	buf.EncodeUint16(m.TTL)
	buf.EncodeUint32(uint32(m.PortIDSubtype))
	buf.EncodeUint32(uint32(m.ChassisIDSubtype))
	return buf.Bytes(), nil
}
func (m *LldpDetails) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.SwIfIndex = interface_types.InterfaceIndex(buf.DecodeUint32())
	m.LastHeard = buf.DecodeFloat64()
	m.LastSent = buf.DecodeFloat64()
	m.ChassisID = make([]byte, 64)
	copy(m.ChassisID, buf.DecodeBytes(len(m.ChassisID)))
	m.ChassisIDLen = buf.DecodeUint8()
	m.PortID = make([]byte, 64)
	copy(m.PortID, buf.DecodeBytes(len(m.PortID)))
	m.PortIDLen = buf.DecodeUint8()
	m.TTL = buf.DecodeUint16()
	m.PortIDSubtype = PortIDSubtype(buf.DecodeUint32())
	m.ChassisIDSubtype = ChassisIDSubtype(buf.DecodeUint32())
	return nil
}

// Dump lldp neighbors
// LldpDump defines message 'lldp_dump'.
type LldpDump struct {
	Cursor uint32 `binapi:"u32,name=cursor" json:"cursor,omitempty"`
}

func (m *LldpDump) Reset()               { *m = LldpDump{} }
func (*LldpDump) GetMessageName() string { return "lldp_dump" }
func (*LldpDump) GetCrcString() string   { return "f75ba505" }
func (*LldpDump) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *LldpDump) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Cursor
	return size
}
func (m *LldpDump) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeUint32(m.Cursor)
	return buf.Bytes(), nil
}
func (m *LldpDump) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Cursor = buf.DecodeUint32()
	return nil
}

// LldpDumpReply defines message 'lldp_dump_reply'.
type LldpDumpReply struct {
	Retval int32  `binapi:"i32,name=retval" json:"retval,omitempty"`
	Cursor uint32 `binapi:"u32,name=cursor" json:"cursor,omitempty"`
}

func (m *LldpDumpReply) Reset()               { *m = LldpDumpReply{} }
func (*LldpDumpReply) GetMessageName() string { return "lldp_dump_reply" }
func (*LldpDumpReply) GetCrcString() string   { return "53b48f5d" }
func (*LldpDumpReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *LldpDumpReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Retval
	size += 4 // m.Cursor
	return size
}
func (m *LldpDumpReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	buf.EncodeUint32(m.Cursor)
	return buf.Bytes(), nil
}
func (m *LldpDumpReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	m.Cursor = buf.DecodeUint32()
	return nil
}

// Interface set LLDP request
//     - sw_if_index - interface for which to enable/disable LLDP
//     - mgmt_ip4_addr - management ip4 address of the interface
//     - mgmt_ip6_addr - management ip6 address of the interface
//     - mgmt_oid - OID(Object Identifier) of the interface
//     - enable - if non-zero enable, else disable
//     - port_desc - local port description
// SwInterfaceSetLldp defines message 'sw_interface_set_lldp'.
type SwInterfaceSetLldp struct {
	SwIfIndex interface_types.InterfaceIndex `binapi:"interface_index,name=sw_if_index" json:"sw_if_index,omitempty"`
	MgmtIP4   ip_types.IP4Address            `binapi:"ip4_address,name=mgmt_ip4" json:"mgmt_ip4,omitempty"`
	MgmtIP6   ip_types.IP6Address            `binapi:"ip6_address,name=mgmt_ip6" json:"mgmt_ip6,omitempty"`
	MgmtOid   []byte                         `binapi:"u8[128],name=mgmt_oid" json:"mgmt_oid,omitempty"`
	Enable    bool                           `binapi:"bool,name=enable,default=true" json:"enable,omitempty"`
	PortDesc  string                         `binapi:"string[],name=port_desc" json:"port_desc,omitempty"`
}

func (m *SwInterfaceSetLldp) Reset()               { *m = SwInterfaceSetLldp{} }
func (*SwInterfaceSetLldp) GetMessageName() string { return "sw_interface_set_lldp" }
func (*SwInterfaceSetLldp) GetCrcString() string   { return "57afbcd4" }
func (*SwInterfaceSetLldp) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *SwInterfaceSetLldp) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4                   // m.SwIfIndex
	size += 1 * 4               // m.MgmtIP4
	size += 1 * 16              // m.MgmtIP6
	size += 1 * 128             // m.MgmtOid
	size += 1                   // m.Enable
	size += 4 + len(m.PortDesc) // m.PortDesc
	return size
}
func (m *SwInterfaceSetLldp) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeUint32(uint32(m.SwIfIndex))
	buf.EncodeBytes(m.MgmtIP4[:], 4)
	buf.EncodeBytes(m.MgmtIP6[:], 16)
	buf.EncodeBytes(m.MgmtOid, 128)
	buf.EncodeBool(m.Enable)
	buf.EncodeString(m.PortDesc, 0)
	return buf.Bytes(), nil
}
func (m *SwInterfaceSetLldp) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.SwIfIndex = interface_types.InterfaceIndex(buf.DecodeUint32())
	copy(m.MgmtIP4[:], buf.DecodeBytes(4))
	copy(m.MgmtIP6[:], buf.DecodeBytes(16))
	m.MgmtOid = make([]byte, 128)
	copy(m.MgmtOid, buf.DecodeBytes(len(m.MgmtOid)))
	m.Enable = buf.DecodeBool()
	m.PortDesc = buf.DecodeString(0)
	return nil
}

// SwInterfaceSetLldpReply defines message 'sw_interface_set_lldp_reply'.
type SwInterfaceSetLldpReply struct {
	Retval int32 `binapi:"i32,name=retval" json:"retval,omitempty"`
}

func (m *SwInterfaceSetLldpReply) Reset()               { *m = SwInterfaceSetLldpReply{} }
func (*SwInterfaceSetLldpReply) GetMessageName() string { return "sw_interface_set_lldp_reply" }
func (*SwInterfaceSetLldpReply) GetCrcString() string   { return "e8d4e804" }
func (*SwInterfaceSetLldpReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *SwInterfaceSetLldpReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Retval
	return size
}
func (m *SwInterfaceSetLldpReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	return buf.Bytes(), nil
}
func (m *SwInterfaceSetLldpReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	return nil
}

func init() { file_lldp_binapi_init() }
func file_lldp_binapi_init() {
	api.RegisterMessage((*LldpConfig)(nil), "lldp_config_c14445df")
	api.RegisterMessage((*LldpConfigReply)(nil), "lldp_config_reply_e8d4e804")
	api.RegisterMessage((*LldpDetails)(nil), "lldp_details_c2d226cd")
	api.RegisterMessage((*LldpDump)(nil), "lldp_dump_f75ba505")
	api.RegisterMessage((*LldpDumpReply)(nil), "lldp_dump_reply_53b48f5d")
	api.RegisterMessage((*SwInterfaceSetLldp)(nil), "sw_interface_set_lldp_57afbcd4")
	api.RegisterMessage((*SwInterfaceSetLldpReply)(nil), "sw_interface_set_lldp_reply_e8d4e804")
}

// Messages returns list of all messages in this module.
func AllMessages() []api.Message {
	return []api.Message{
		(*LldpConfig)(nil),
		(*LldpConfigReply)(nil),
		(*LldpDetails)(nil),
		(*LldpDump)(nil),
		(*LldpDumpReply)(nil),
		(*SwInterfaceSetLldp)(nil),
		(*SwInterfaceSetLldpReply)(nil),
	}
}