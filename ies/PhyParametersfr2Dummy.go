package ies

import "rrc/utils"

// Phy-ParametersFR2-dummy ::= ENUMERATED
type PhyParametersfr2Dummy struct {
	Value utils.ENUMERATED
}

const (
	PhyParametersfr2DummyEnumeratedNothing = iota
	PhyParametersfr2DummyEnumeratedSupported
)
