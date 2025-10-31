package ies

import "rrc/utils"

// Phy-ParametersFR1-scs-60kHz ::= ENUMERATED
type PhyParametersfr1Scs60khz struct {
	Value utils.ENUMERATED
}

const (
	PhyParametersfr1Scs60khzEnumeratedNothing = iota
	PhyParametersfr1Scs60khzEnumeratedSupported
)
