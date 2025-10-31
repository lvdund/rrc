package ies

import "rrc/utils"

// NR-RS-Type ::= ENUMERATED
type NrRsType struct {
	Value utils.ENUMERATED
}

const (
	NrRsTypeEnumeratedNothing = iota
	NrRsTypeEnumeratedSsb
	NrRsTypeEnumeratedCsi_Rs
)
