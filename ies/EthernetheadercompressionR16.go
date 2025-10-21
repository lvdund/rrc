package ies

import "rrc/utils"

// EthernetHeaderCompression-r16 ::= SEQUENCE
// Extensible
type EthernetheadercompressionR16 struct {
	EhcCommonR16   interface{}
	EhcDownlinkR16 *interface{}
	EhcUplinkR16   *interface{}
}
