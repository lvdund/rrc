package ies

import "rrc/utils"

// PDCP-Parameters-non-DRB-IAB-r16 ::= ENUMERATED
type PdcpParametersNonDrbIabR16 struct {
	Value utils.ENUMERATED
}

const (
	PdcpParametersNonDrbIabR16EnumeratedNothing = iota
	PdcpParametersNonDrbIabR16EnumeratedSupported
)
