// Code generated by GoVPP's binapi-generator. DO NOT EDIT.
// versions:
//  binapi-generator: v0.10.0
//  VPP:              24.10-rc0~84-gd2d41bc34
// source: plugins/testing.api.json

// Package testing contains generated bindings for API file testing.api.
//
// Contents:
// -  2 messages
package testing

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
	APIFile    = "testing"
	APIVersion = "0.1.0"
	VersionCrc = 0x3a1a2c50
)

// @brief API to enable / disable testing on an interface
//     - enable_disable - 1 to enable, 0 to disable the feature
//     - sw_if_index - interface handle
// TestingEnableDisable defines message 'testing_enable_disable'.
type TestingEnableDisable struct {
	EnableDisable bool                           `binapi:"bool,name=enable_disable" json:"enable_disable,omitempty"`
	SwIfIndex     interface_types.InterfaceIndex `binapi:"interface_index,name=sw_if_index" json:"sw_if_index,omitempty"`
}

func (m *TestingEnableDisable) Reset()               { *m = TestingEnableDisable{} }
func (*TestingEnableDisable) GetMessageName() string { return "testing_enable_disable" }
func (*TestingEnableDisable) GetCrcString() string   { return "3865946c" }
func (*TestingEnableDisable) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *TestingEnableDisable) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 1 // m.EnableDisable
	size += 4 // m.SwIfIndex
	return size
}
func (m *TestingEnableDisable) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeBool(m.EnableDisable)
	buf.EncodeUint32(uint32(m.SwIfIndex))
	return buf.Bytes(), nil
}
func (m *TestingEnableDisable) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.EnableDisable = buf.DecodeBool()
	m.SwIfIndex = interface_types.InterfaceIndex(buf.DecodeUint32())
	return nil
}

// TestingEnableDisableReply defines message 'testing_enable_disable_reply'.
type TestingEnableDisableReply struct {
	Retval int32 `binapi:"i32,name=retval" json:"retval,omitempty"`
}

func (m *TestingEnableDisableReply) Reset()               { *m = TestingEnableDisableReply{} }
func (*TestingEnableDisableReply) GetMessageName() string { return "testing_enable_disable_reply" }
func (*TestingEnableDisableReply) GetCrcString() string   { return "e8d4e804" }
func (*TestingEnableDisableReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *TestingEnableDisableReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Retval
	return size
}
func (m *TestingEnableDisableReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	return buf.Bytes(), nil
}
func (m *TestingEnableDisableReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	return nil
}

func init() { file_testing_binapi_init() }
func file_testing_binapi_init() {
	api.RegisterMessage((*TestingEnableDisable)(nil), "testing_enable_disable_3865946c")
	api.RegisterMessage((*TestingEnableDisableReply)(nil), "testing_enable_disable_reply_e8d4e804")
}

// Messages returns list of all messages in this module.
func AllMessages() []api.Message {
	return []api.Message{
		(*TestingEnableDisable)(nil),
		(*TestingEnableDisableReply)(nil),
	}
}
