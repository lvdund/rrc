package ies

import "rrc/utils"

// MIMO-CapabilityUL-r10 ::= ENUMERATED
type MimoCapabilityulR10 struct {
	Value utils.ENUMERATED
}

const (
	MimoCapabilityulR10Twolayers  = 0
	MimoCapabilityulR10Fourlayers = 1
)
