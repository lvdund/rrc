package ies

import "rrc/utils"

// RLC-Parameters-NB-r15-rlc-UM-r15 ::= ENUMERATED
type RlcParametersNbR15RlcUmR15 struct {
	Value utils.ENUMERATED
}

const (
	RlcParametersNbR15RlcUmR15EnumeratedNothing = iota
	RlcParametersNbR15RlcUmR15EnumeratedSupported
)
