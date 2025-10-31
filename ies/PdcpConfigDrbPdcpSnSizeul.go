package ies

import "rrc/utils"

// PDCP-Config-drb-pdcp-SN-SizeUL ::= ENUMERATED
type PdcpConfigDrbPdcpSnSizeul struct {
	Value utils.ENUMERATED
}

const (
	PdcpConfigDrbPdcpSnSizeulEnumeratedNothing = iota
	PdcpConfigDrbPdcpSnSizeulEnumeratedLen12bits
	PdcpConfigDrbPdcpSnSizeulEnumeratedLen18bits
)
