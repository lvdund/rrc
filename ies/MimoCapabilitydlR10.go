package ies

import "rrc/utils"

// MIMO-CapabilityDL-r10 ::= ENUMERATED
type MimoCapabilitydlR10 struct {
	Value utils.ENUMERATED
}

const (
	MimoCapabilitydlR10Twolayers   = 0
	MimoCapabilitydlR10Fourlayers  = 1
	MimoCapabilitydlR10Eightlayers = 2
)
