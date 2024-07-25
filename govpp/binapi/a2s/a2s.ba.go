// Code generated by GoVPP's binapi-generator. DO NOT EDIT.
// versions:
//  binapi-generator: v0.10.0
//  VPP:              24.10-rc0~84-gd2d41bc34
// source: plugins/a2s.api.json

// Package a2s contains generated bindings for API file a2s.api.
//
// Contents:
// -  3 structs
// -  4 messages
package a2s

import (
	"govpp/binapi/interface_types"

	api "go.fd.io/govpp/api"
	codec "go.fd.io/govpp/codec"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the GoVPP api package it is being compiled against.
// A compilation error at this line likely means your copy of the
// GoVPP api package needs to be updated.
const _ = api.GoVppAPIPackageIsVersion2

const (
	APIFile    = "a2s"
	APIVersion = "0.2.0"
	VersionCrc = 0xbf6f47ab
)

// A2sData defines type 'a2s_data'.
type A2sData struct {
	InfoData     []byte `binapi:"u8[2000],name=INFO_DATA" json:"INFO_DATA,omitempty"`
	PlayerData   []byte `binapi:"u8[2000],name=PLAYER_DATA" json:"PLAYER_DATA,omitempty"`
	RulesData    []byte `binapi:"u8[10000],name=RULES_DATA" json:"RULES_DATA,omitempty"`
	InfoLength   uint16 `binapi:"u16,name=info_length" json:"info_length,omitempty"`
	PlayerLength uint16 `binapi:"u16,name=player_length" json:"player_length,omitempty"`
	RulesLength  uint16 `binapi:"u16,name=rules_length" json:"rules_length,omitempty"`
	Strict       bool   `binapi:"bool,name=strict" json:"strict,omitempty"`
}

// A2sKey defines type 'a2s_key'.
type A2sKey struct {
	IP   uint32 `binapi:"u32,name=ip" json:"ip,omitempty"`
	Port uint16 `binapi:"u16,name=port" json:"port,omitempty"`
}

// ClibBihashKvA2s defines type 'clib_bihash_kv_a2s'.
type ClibBihashKvA2s struct {
	Key   A2sKey  `binapi:"a2s_key,name=key" json:"key,omitempty"`
	Value A2sData `binapi:"a2s_data,name=value" json:"value,omitempty"`
}

// @brief API to enable / disable a2s on an interface
//     - enable_disable - 1 to enable, 0 to disable the feature
//     - sw_if_index - interface handle
// A2sEnableDisable defines message 'a2s_enable_disable'.
type A2sEnableDisable struct {
	EnableDisable bool                           `binapi:"bool,name=enable_disable" json:"enable_disable,omitempty"`
	SwIfIndex     interface_types.InterfaceIndex `binapi:"interface_index,name=sw_if_index" json:"sw_if_index,omitempty"`
}

func (m *A2sEnableDisable) Reset()               { *m = A2sEnableDisable{} }
func (*A2sEnableDisable) GetMessageName() string { return "a2s_enable_disable" }
func (*A2sEnableDisable) GetCrcString() string   { return "3865946c" }
func (*A2sEnableDisable) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *A2sEnableDisable) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 1 // m.EnableDisable
	size += 4 // m.SwIfIndex
	return size
}
func (m *A2sEnableDisable) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeBool(m.EnableDisable)
	buf.EncodeUint32(uint32(m.SwIfIndex))
	return buf.Bytes(), nil
}
func (m *A2sEnableDisable) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.EnableDisable = buf.DecodeBool()
	m.SwIfIndex = interface_types.InterfaceIndex(buf.DecodeUint32())
	return nil
}

// A2sEnableDisableReply defines message 'a2s_enable_disable_reply'.
type A2sEnableDisableReply struct {
	Retval int32 `binapi:"i32,name=retval" json:"retval,omitempty"`
}

func (m *A2sEnableDisableReply) Reset()               { *m = A2sEnableDisableReply{} }
func (*A2sEnableDisableReply) GetMessageName() string { return "a2s_enable_disable_reply" }
func (*A2sEnableDisableReply) GetCrcString() string   { return "e8d4e804" }
func (*A2sEnableDisableReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *A2sEnableDisableReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Retval
	return size
}
func (m *A2sEnableDisableReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	return buf.Bytes(), nil
}
func (m *A2sEnableDisableReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	return nil
}

// @brief API to set a2s data
//     - is_add - Boolean for adding or removing
//     - kv - Key value pair to be inserted or removed
// A2sSetData defines message 'a2s_set_data'.
type A2sSetData struct {
	IsAdd bool            `binapi:"bool,name=is_add" json:"is_add,omitempty"`
	Kv    ClibBihashKvA2s `binapi:"clib_bihash_kv_a2s,name=kv" json:"kv,omitempty"`
}

func (m *A2sSetData) Reset()               { *m = A2sSetData{} }
func (*A2sSetData) GetMessageName() string { return "a2s_set_data" }
func (*A2sSetData) GetCrcString() string   { return "dc3b0644" }
func (*A2sSetData) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *A2sSetData) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 1         // m.IsAdd
	size += 4         // m.Kv.Key.IP
	size += 2         // m.Kv.Key.Port
	size += 1 * 2000  // m.Kv.Value.InfoData
	size += 1 * 2000  // m.Kv.Value.PlayerData
	size += 1 * 10000 // m.Kv.Value.RulesData
	size += 2         // m.Kv.Value.InfoLength
	size += 2         // m.Kv.Value.PlayerLength
	size += 2         // m.Kv.Value.RulesLength
	size += 1         // m.Kv.Value.Strict
	return size
}
func (m *A2sSetData) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeBool(m.IsAdd)
	buf.EncodeUint32(m.Kv.Key.IP)
	buf.EncodeUint16(m.Kv.Key.Port)
	buf.EncodeBytes(m.Kv.Value.InfoData, 2000)
	buf.EncodeBytes(m.Kv.Value.PlayerData, 2000)
	buf.EncodeBytes(m.Kv.Value.RulesData, 10000)
	buf.EncodeUint16(m.Kv.Value.InfoLength)
	buf.EncodeUint16(m.Kv.Value.PlayerLength)
	buf.EncodeUint16(m.Kv.Value.RulesLength)
	buf.EncodeBool(m.Kv.Value.Strict)
	return buf.Bytes(), nil
}
func (m *A2sSetData) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.IsAdd = buf.DecodeBool()
	m.Kv.Key.IP = buf.DecodeUint32()
	m.Kv.Key.Port = buf.DecodeUint16()
	m.Kv.Value.InfoData = make([]byte, 2000)
	copy(m.Kv.Value.InfoData, buf.DecodeBytes(len(m.Kv.Value.InfoData)))
	m.Kv.Value.PlayerData = make([]byte, 2000)
	copy(m.Kv.Value.PlayerData, buf.DecodeBytes(len(m.Kv.Value.PlayerData)))
	m.Kv.Value.RulesData = make([]byte, 10000)
	copy(m.Kv.Value.RulesData, buf.DecodeBytes(len(m.Kv.Value.RulesData)))
	m.Kv.Value.InfoLength = buf.DecodeUint16()
	m.Kv.Value.PlayerLength = buf.DecodeUint16()
	m.Kv.Value.RulesLength = buf.DecodeUint16()
	m.Kv.Value.Strict = buf.DecodeBool()
	return nil
}

// A2sSetDataReply defines message 'a2s_set_data_reply'.
type A2sSetDataReply struct {
	Retval int32 `binapi:"i32,name=retval" json:"retval,omitempty"`
}

func (m *A2sSetDataReply) Reset()               { *m = A2sSetDataReply{} }
func (*A2sSetDataReply) GetMessageName() string { return "a2s_set_data_reply" }
func (*A2sSetDataReply) GetCrcString() string   { return "e8d4e804" }
func (*A2sSetDataReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *A2sSetDataReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Retval
	return size
}
func (m *A2sSetDataReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	return buf.Bytes(), nil
}
func (m *A2sSetDataReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	return nil
}

func init() { file_a2s_binapi_init() }
func file_a2s_binapi_init() {
	api.RegisterMessage((*A2sEnableDisable)(nil), "a2s_enable_disable_3865946c")
	api.RegisterMessage((*A2sEnableDisableReply)(nil), "a2s_enable_disable_reply_e8d4e804")
	api.RegisterMessage((*A2sSetData)(nil), "a2s_set_data_dc3b0644")
	api.RegisterMessage((*A2sSetDataReply)(nil), "a2s_set_data_reply_e8d4e804")
}

// Messages returns list of all messages in this module.
func AllMessages() []api.Message {
	return []api.Message{
		(*A2sEnableDisable)(nil),
		(*A2sEnableDisableReply)(nil),
		(*A2sSetData)(nil),
		(*A2sSetDataReply)(nil),
	}
}
