package ies

import "rrc/utils"

// Phy-ParametersCommon-dummy ::= ENUMERATED
type PhyParameterscommonDummy struct {
	Value utils.ENUMERATED
}

const (
	PhyParameterscommonDummyEnumeratedNothing = iota
	PhyParameterscommonDummyEnumeratedSupported
)
