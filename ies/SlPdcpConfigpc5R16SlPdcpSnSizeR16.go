package ies

import "rrc/utils"

// SL-PDCP-ConfigPC5-r16-sl-PDCP-SN-Size-r16 ::= ENUMERATED
type SlPdcpConfigpc5R16SlPdcpSnSizeR16 struct {
	Value utils.ENUMERATED
}

const (
	SlPdcpConfigpc5R16SlPdcpSnSizeR16EnumeratedNothing = iota
	SlPdcpConfigpc5R16SlPdcpSnSizeR16EnumeratedLen12bits
	SlPdcpConfigpc5R16SlPdcpSnSizeR16EnumeratedLen18bits
)
