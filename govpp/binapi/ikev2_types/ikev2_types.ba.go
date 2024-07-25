// Code generated by GoVPP's binapi-generator. DO NOT EDIT.
// versions:
//  binapi-generator: v0.10.0
//  VPP:              24.10-rc0~84-gd2d41bc34
// source: plugins/ikev2_types.api.json

// Package ikev2_types contains generated bindings for API file ikev2_types.api.
//
// Contents:
// -  1 enum
// - 15 structs
package ikev2_types

import (
	"govpp/binapi/interface_types"
	"govpp/binapi/ip_types"
	"strconv"

	api "go.fd.io/govpp/api"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the GoVPP api package it is being compiled against.
// A compilation error at this line likely means your copy of the
// GoVPP api package needs to be updated.
const _ = api.GoVppAPIPackageIsVersion2

const (
	APIFile    = "ikev2_types"
	APIVersion = "1.0.0"
	VersionCrc = 0x64c72418
)

// Ikev2State defines enum 'ikev2_state'.
type Ikev2State uint32

const (
	UNKNOWN            Ikev2State = 0
	SA_INIT            Ikev2State = 1
	DELETED            Ikev2State = 2
	AUTH_FAILED        Ikev2State = 3
	AUTHENTICATED      Ikev2State = 4
	NOTIFY_AND_DELETE  Ikev2State = 5
	TS_UNACCEPTABLE    Ikev2State = 6
	NO_PROPOSAL_CHOSEN Ikev2State = 7
)

var (
	Ikev2State_name = map[uint32]string{
		0: "UNKNOWN",
		1: "SA_INIT",
		2: "DELETED",
		3: "AUTH_FAILED",
		4: "AUTHENTICATED",
		5: "NOTIFY_AND_DELETE",
		6: "TS_UNACCEPTABLE",
		7: "NO_PROPOSAL_CHOSEN",
	}
	Ikev2State_value = map[string]uint32{
		"UNKNOWN":            0,
		"SA_INIT":            1,
		"DELETED":            2,
		"AUTH_FAILED":        3,
		"AUTHENTICATED":      4,
		"NOTIFY_AND_DELETE":  5,
		"TS_UNACCEPTABLE":    6,
		"NO_PROPOSAL_CHOSEN": 7,
	}
)

func (x Ikev2State) String() string {
	s, ok := Ikev2State_name[uint32(x)]
	if ok {
		return s
	}
	return "Ikev2State(" + strconv.Itoa(int(x)) + ")"
}

// Ikev2Auth defines type 'ikev2_auth'.
type Ikev2Auth struct {
	Method  uint8  `binapi:"u8,name=method" json:"method,omitempty"`
	Hex     uint8  `binapi:"u8,name=hex" json:"hex,omitempty"`
	DataLen uint32 `binapi:"u32,name=data_len" json:"-"`
	Data    []byte `binapi:"u8[data_len],name=data" json:"data,omitempty"`
}

// Ikev2ChildSa defines type 'ikev2_child_sa'.
type Ikev2ChildSa struct {
	SaIndex      uint32           `binapi:"u32,name=sa_index" json:"sa_index,omitempty"`
	ChildSaIndex uint32           `binapi:"u32,name=child_sa_index" json:"child_sa_index,omitempty"`
	ISpi         uint32           `binapi:"u32,name=i_spi" json:"i_spi,omitempty"`
	RSpi         uint32           `binapi:"u32,name=r_spi" json:"r_spi,omitempty"`
	Keys         Ikev2Keys        `binapi:"ikev2_keys,name=keys" json:"keys,omitempty"`
	Encryption   Ikev2SaTransform `binapi:"ikev2_sa_transform,name=encryption" json:"encryption,omitempty"`
	Integrity    Ikev2SaTransform `binapi:"ikev2_sa_transform,name=integrity" json:"integrity,omitempty"`
	Esn          Ikev2SaTransform `binapi:"ikev2_sa_transform,name=esn" json:"esn,omitempty"`
}

// Ikev2ChildSaV2 defines type 'ikev2_child_sa_v2'.
type Ikev2ChildSaV2 struct {
	SaIndex      uint32           `binapi:"u32,name=sa_index" json:"sa_index,omitempty"`
	ChildSaIndex uint32           `binapi:"u32,name=child_sa_index" json:"child_sa_index,omitempty"`
	ISpi         uint32           `binapi:"u32,name=i_spi" json:"i_spi,omitempty"`
	RSpi         uint32           `binapi:"u32,name=r_spi" json:"r_spi,omitempty"`
	Keys         Ikev2Keys        `binapi:"ikev2_keys,name=keys" json:"keys,omitempty"`
	Encryption   Ikev2SaTransform `binapi:"ikev2_sa_transform,name=encryption" json:"encryption,omitempty"`
	Integrity    Ikev2SaTransform `binapi:"ikev2_sa_transform,name=integrity" json:"integrity,omitempty"`
	Esn          Ikev2SaTransform `binapi:"ikev2_sa_transform,name=esn" json:"esn,omitempty"`
	Uptime       float64          `binapi:"f64,name=uptime" json:"uptime,omitempty"`
}

// Ikev2EspTransforms defines type 'ikev2_esp_transforms'.
type Ikev2EspTransforms struct {
	CryptoAlg     uint8  `binapi:"u8,name=crypto_alg" json:"crypto_alg,omitempty"`
	CryptoKeySize uint32 `binapi:"u32,name=crypto_key_size" json:"crypto_key_size,omitempty"`
	IntegAlg      uint8  `binapi:"u8,name=integ_alg" json:"integ_alg,omitempty"`
}

// Ikev2ID defines type 'ikev2_id'.
type Ikev2ID struct {
	Type    uint8  `binapi:"u8,name=type" json:"type,omitempty"`
	DataLen uint8  `binapi:"u8,name=data_len" json:"data_len,omitempty"`
	Data    string `binapi:"string[64],name=data" json:"data,omitempty"`
}

// Ikev2IkeTransforms defines type 'ikev2_ike_transforms'.
type Ikev2IkeTransforms struct {
	CryptoAlg     uint8  `binapi:"u8,name=crypto_alg" json:"crypto_alg,omitempty"`
	CryptoKeySize uint32 `binapi:"u32,name=crypto_key_size" json:"crypto_key_size,omitempty"`
	IntegAlg      uint8  `binapi:"u8,name=integ_alg" json:"integ_alg,omitempty"`
	DhGroup       uint8  `binapi:"u8,name=dh_group" json:"dh_group,omitempty"`
}

// Ikev2Keys defines type 'ikev2_keys'.
type Ikev2Keys struct {
	SkD     []byte `binapi:"u8[64],name=sk_d" json:"sk_d,omitempty"`
	SkDLen  uint8  `binapi:"u8,name=sk_d_len" json:"sk_d_len,omitempty"`
	SkAi    []byte `binapi:"u8[64],name=sk_ai" json:"sk_ai,omitempty"`
	SkAiLen uint8  `binapi:"u8,name=sk_ai_len" json:"sk_ai_len,omitempty"`
	SkAr    []byte `binapi:"u8[64],name=sk_ar" json:"sk_ar,omitempty"`
	SkArLen uint8  `binapi:"u8,name=sk_ar_len" json:"sk_ar_len,omitempty"`
	SkEi    []byte `binapi:"u8[64],name=sk_ei" json:"sk_ei,omitempty"`
	SkEiLen uint8  `binapi:"u8,name=sk_ei_len" json:"sk_ei_len,omitempty"`
	SkEr    []byte `binapi:"u8[64],name=sk_er" json:"sk_er,omitempty"`
	SkErLen uint8  `binapi:"u8,name=sk_er_len" json:"sk_er_len,omitempty"`
	SkPi    []byte `binapi:"u8[64],name=sk_pi" json:"sk_pi,omitempty"`
	SkPiLen uint8  `binapi:"u8,name=sk_pi_len" json:"sk_pi_len,omitempty"`
	SkPr    []byte `binapi:"u8[64],name=sk_pr" json:"sk_pr,omitempty"`
	SkPrLen uint8  `binapi:"u8,name=sk_pr_len" json:"sk_pr_len,omitempty"`
}

// Ikev2Profile defines type 'ikev2_profile'.
type Ikev2Profile struct {
	Name             string             `binapi:"string[64],name=name" json:"name,omitempty"`
	LocID            Ikev2ID            `binapi:"ikev2_id,name=loc_id" json:"loc_id,omitempty"`
	RemID            Ikev2ID            `binapi:"ikev2_id,name=rem_id" json:"rem_id,omitempty"`
	LocTs            Ikev2Ts            `binapi:"ikev2_ts,name=loc_ts" json:"loc_ts,omitempty"`
	RemTs            Ikev2Ts            `binapi:"ikev2_ts,name=rem_ts" json:"rem_ts,omitempty"`
	Responder        Ikev2Responder     `binapi:"ikev2_responder,name=responder" json:"responder,omitempty"`
	IkeTs            Ikev2IkeTransforms `binapi:"ikev2_ike_transforms,name=ike_ts" json:"ike_ts,omitempty"`
	EspTs            Ikev2EspTransforms `binapi:"ikev2_esp_transforms,name=esp_ts" json:"esp_ts,omitempty"`
	Lifetime         uint64             `binapi:"u64,name=lifetime" json:"lifetime,omitempty"`
	LifetimeMaxdata  uint64             `binapi:"u64,name=lifetime_maxdata" json:"lifetime_maxdata,omitempty"`
	LifetimeJitter   uint32             `binapi:"u32,name=lifetime_jitter" json:"lifetime_jitter,omitempty"`
	Handover         uint32             `binapi:"u32,name=handover" json:"handover,omitempty"`
	IpsecOverUDPPort uint16             `binapi:"u16,name=ipsec_over_udp_port" json:"ipsec_over_udp_port,omitempty"`
	TunItf           uint32             `binapi:"u32,name=tun_itf" json:"tun_itf,omitempty"`
	UDPEncap         bool               `binapi:"bool,name=udp_encap" json:"udp_encap,omitempty"`
	NattDisabled     bool               `binapi:"bool,name=natt_disabled" json:"natt_disabled,omitempty"`
	Auth             Ikev2Auth          `binapi:"ikev2_auth,name=auth" json:"auth,omitempty"`
}

// Ikev2Responder defines type 'ikev2_responder'.
type Ikev2Responder struct {
	SwIfIndex interface_types.InterfaceIndex `binapi:"interface_index,name=sw_if_index" json:"sw_if_index,omitempty"`
	Addr      ip_types.Address               `binapi:"address,name=addr" json:"addr,omitempty"`
}

// Ikev2Sa defines type 'ikev2_sa'.
type Ikev2Sa struct {
	SaIndex      uint32           `binapi:"u32,name=sa_index" json:"sa_index,omitempty"`
	ProfileIndex uint32           `binapi:"u32,name=profile_index" json:"profile_index,omitempty"`
	Ispi         uint64           `binapi:"u64,name=ispi" json:"ispi,omitempty"`
	Rspi         uint64           `binapi:"u64,name=rspi" json:"rspi,omitempty"`
	Iaddr        ip_types.Address `binapi:"address,name=iaddr" json:"iaddr,omitempty"`
	Raddr        ip_types.Address `binapi:"address,name=raddr" json:"raddr,omitempty"`
	Keys         Ikev2Keys        `binapi:"ikev2_keys,name=keys" json:"keys,omitempty"`
	IID          Ikev2ID          `binapi:"ikev2_id,name=i_id" json:"i_id,omitempty"`
	RID          Ikev2ID          `binapi:"ikev2_id,name=r_id" json:"r_id,omitempty"`
	Encryption   Ikev2SaTransform `binapi:"ikev2_sa_transform,name=encryption" json:"encryption,omitempty"`
	Integrity    Ikev2SaTransform `binapi:"ikev2_sa_transform,name=integrity" json:"integrity,omitempty"`
	Prf          Ikev2SaTransform `binapi:"ikev2_sa_transform,name=prf" json:"prf,omitempty"`
	Dh           Ikev2SaTransform `binapi:"ikev2_sa_transform,name=dh" json:"dh,omitempty"`
	Stats        Ikev2SaStats     `binapi:"ikev2_sa_stats,name=stats" json:"stats,omitempty"`
}

// Ikev2SaStats defines type 'ikev2_sa_stats'.
type Ikev2SaStats struct {
	NKeepalives       uint16 `binapi:"u16,name=n_keepalives" json:"n_keepalives,omitempty"`
	NRekeyReq         uint16 `binapi:"u16,name=n_rekey_req" json:"n_rekey_req,omitempty"`
	NSaInitReq        uint16 `binapi:"u16,name=n_sa_init_req" json:"n_sa_init_req,omitempty"`
	NSaAuthReq        uint16 `binapi:"u16,name=n_sa_auth_req" json:"n_sa_auth_req,omitempty"`
	NRetransmit       uint16 `binapi:"u16,name=n_retransmit" json:"n_retransmit,omitempty"`
	NInitSaRetransmit uint16 `binapi:"u16,name=n_init_sa_retransmit" json:"n_init_sa_retransmit,omitempty"`
}

// Ikev2SaTransform defines type 'ikev2_sa_transform'.
type Ikev2SaTransform struct {
	TransformType uint8  `binapi:"u8,name=transform_type" json:"transform_type,omitempty"`
	TransformID   uint16 `binapi:"u16,name=transform_id" json:"transform_id,omitempty"`
	KeyLen        uint16 `binapi:"u16,name=key_len" json:"key_len,omitempty"`
	KeyTrunc      uint16 `binapi:"u16,name=key_trunc" json:"key_trunc,omitempty"`
	BlockSize     uint16 `binapi:"u16,name=block_size" json:"block_size,omitempty"`
	DhGroup       uint8  `binapi:"u8,name=dh_group" json:"dh_group,omitempty"`
}

// Ikev2SaV2 defines type 'ikev2_sa_v2'.
type Ikev2SaV2 struct {
	SaIndex     uint32           `binapi:"u32,name=sa_index" json:"sa_index,omitempty"`
	ProfileName string           `binapi:"string[64],name=profile_name" json:"profile_name,omitempty"`
	State       Ikev2State       `binapi:"ikev2_state,name=state" json:"state,omitempty"`
	Ispi        uint64           `binapi:"u64,name=ispi" json:"ispi,omitempty"`
	Rspi        uint64           `binapi:"u64,name=rspi" json:"rspi,omitempty"`
	Iaddr       ip_types.Address `binapi:"address,name=iaddr" json:"iaddr,omitempty"`
	Raddr       ip_types.Address `binapi:"address,name=raddr" json:"raddr,omitempty"`
	Keys        Ikev2Keys        `binapi:"ikev2_keys,name=keys" json:"keys,omitempty"`
	IID         Ikev2ID          `binapi:"ikev2_id,name=i_id" json:"i_id,omitempty"`
	RID         Ikev2ID          `binapi:"ikev2_id,name=r_id" json:"r_id,omitempty"`
	Encryption  Ikev2SaTransform `binapi:"ikev2_sa_transform,name=encryption" json:"encryption,omitempty"`
	Integrity   Ikev2SaTransform `binapi:"ikev2_sa_transform,name=integrity" json:"integrity,omitempty"`
	Prf         Ikev2SaTransform `binapi:"ikev2_sa_transform,name=prf" json:"prf,omitempty"`
	Dh          Ikev2SaTransform `binapi:"ikev2_sa_transform,name=dh" json:"dh,omitempty"`
	Stats       Ikev2SaStats     `binapi:"ikev2_sa_stats,name=stats" json:"stats,omitempty"`
}

// Ikev2SaV3 defines type 'ikev2_sa_v3'.
type Ikev2SaV3 struct {
	SaIndex     uint32           `binapi:"u32,name=sa_index" json:"sa_index,omitempty"`
	ProfileName string           `binapi:"string[64],name=profile_name" json:"profile_name,omitempty"`
	State       Ikev2State       `binapi:"ikev2_state,name=state" json:"state,omitempty"`
	Ispi        uint64           `binapi:"u64,name=ispi" json:"ispi,omitempty"`
	Rspi        uint64           `binapi:"u64,name=rspi" json:"rspi,omitempty"`
	Iaddr       ip_types.Address `binapi:"address,name=iaddr" json:"iaddr,omitempty"`
	Raddr       ip_types.Address `binapi:"address,name=raddr" json:"raddr,omitempty"`
	Keys        Ikev2Keys        `binapi:"ikev2_keys,name=keys" json:"keys,omitempty"`
	IID         Ikev2ID          `binapi:"ikev2_id,name=i_id" json:"i_id,omitempty"`
	RID         Ikev2ID          `binapi:"ikev2_id,name=r_id" json:"r_id,omitempty"`
	Encryption  Ikev2SaTransform `binapi:"ikev2_sa_transform,name=encryption" json:"encryption,omitempty"`
	Integrity   Ikev2SaTransform `binapi:"ikev2_sa_transform,name=integrity" json:"integrity,omitempty"`
	Prf         Ikev2SaTransform `binapi:"ikev2_sa_transform,name=prf" json:"prf,omitempty"`
	Dh          Ikev2SaTransform `binapi:"ikev2_sa_transform,name=dh" json:"dh,omitempty"`
	Stats       Ikev2SaStats     `binapi:"ikev2_sa_stats,name=stats" json:"stats,omitempty"`
	Uptime      float64          `binapi:"f64,name=uptime" json:"uptime,omitempty"`
}

// Ikev2Ts defines type 'ikev2_ts'.
type Ikev2Ts struct {
	SaIndex      uint32           `binapi:"u32,name=sa_index" json:"sa_index,omitempty"`
	ChildSaIndex uint32           `binapi:"u32,name=child_sa_index" json:"child_sa_index,omitempty"`
	IsLocal      bool             `binapi:"bool,name=is_local" json:"is_local,omitempty"`
	ProtocolID   uint8            `binapi:"u8,name=protocol_id" json:"protocol_id,omitempty"`
	StartPort    uint16           `binapi:"u16,name=start_port" json:"start_port,omitempty"`
	EndPort      uint16           `binapi:"u16,name=end_port" json:"end_port,omitempty"`
	StartAddr    ip_types.Address `binapi:"address,name=start_addr" json:"start_addr,omitempty"`
	EndAddr      ip_types.Address `binapi:"address,name=end_addr" json:"end_addr,omitempty"`
}
