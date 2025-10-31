package ies

import "rrc/utils"

// SL-PDCP-Config-r16-sl-PDCP-SN-Size-r16 ::= ENUMERATED
type SlPdcpConfigR16SlPdcpSnSizeR16 struct {
	Value utils.ENUMERATED
}

const (
	SlPdcpConfigR16SlPdcpSnSizeR16EnumeratedNothing = iota
	SlPdcpConfigR16SlPdcpSnSizeR16EnumeratedLen12bits
	SlPdcpConfigR16SlPdcpSnSizeR16EnumeratedLen18bits
)
