package ies

import "rrc/utils"

// PDCP-Parameters-drb-IAB-r16 ::= ENUMERATED
type PdcpParametersDrbIabR16 struct {
	Value utils.ENUMERATED
}

const (
	PdcpParametersDrbIabR16EnumeratedNothing = iota
	PdcpParametersDrbIabR16EnumeratedSupported
)
