package mhfpacket

import (
	"github.com/Andoryuuta/Erupe/network"
	"github.com/Andoryuuta/byteframe"
)

// MsgSysDeleteObject represents the MSG_SYS_DELETE_OBJECT
type MsgSysDeleteObject struct{}

// Opcode returns the ID associated with this packet type.
func (m *MsgSysDeleteObject) Opcode() network.PacketID {
	return network.MSG_SYS_DELETE_OBJECT
}

// Parse parses the packet from binary
func (m *MsgSysDeleteObject) Parse(bf *byteframe.ByteFrame) error {
	panic("Not implemented")
}

// Build builds a binary packet from the current data.
func (m *MsgSysDeleteObject) Build(bf *byteframe.ByteFrame) error {
	panic("Not implemented")
}