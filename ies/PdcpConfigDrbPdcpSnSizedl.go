package ies

import "rrc/utils"

// PDCP-Config-drb-pdcp-SN-SizeDL ::= ENUMERATED
type PdcpConfigDrbPdcpSnSizedl struct {
	Value utils.ENUMERATED
}

const (
	PdcpConfigDrbPdcpSnSizedlEnumeratedNothing = iota
	PdcpConfigDrbPdcpSnSizedlEnumeratedLen12bits
	PdcpConfigDrbPdcpSnSizedlEnumeratedLen18bits
)
