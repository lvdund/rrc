package ies

import "rrc/utils"

// EthernetHeaderCompression-r16-ehc-Common-r16-ehc-CID-Length-r16 ::= ENUMERATED
type EthernetheadercompressionR16EhcCommonR16EhcCidLengthR16 struct {
	Value utils.ENUMERATED
}

const (
	EthernetheadercompressionR16EhcCommonR16EhcCidLengthR16EnumeratedNothing = iota
	EthernetheadercompressionR16EhcCommonR16EhcCidLengthR16EnumeratedBits7
	EthernetheadercompressionR16EhcCommonR16EhcCidLengthR16EnumeratedBits15
)
