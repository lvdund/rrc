package ies

import "rrc/utils"

// PDCP-Config-rlc-UM-pdcp-SN-Size ::= ENUMERATED
type PdcpConfigRlcUmPdcpSnSize struct {
	Value utils.ENUMERATED
}

const (
	PdcpConfigRlcUmPdcpSnSizeEnumeratedNothing = iota
	PdcpConfigRlcUmPdcpSnSizeEnumeratedLen7bits
	PdcpConfigRlcUmPdcpSnSizeEnumeratedLen12bits
)
