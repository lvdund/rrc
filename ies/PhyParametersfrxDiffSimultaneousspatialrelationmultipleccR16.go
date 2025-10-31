package ies

import "rrc/utils"

// Phy-ParametersFRX-Diff-simultaneousSpatialRelationMultipleCC-r16 ::= ENUMERATED
type PhyParametersfrxDiffSimultaneousspatialrelationmultipleccR16 struct {
	Value utils.ENUMERATED
}

const (
	PhyParametersfrxDiffSimultaneousspatialrelationmultipleccR16EnumeratedNothing = iota
	PhyParametersfrxDiffSimultaneousspatialrelationmultipleccR16EnumeratedSupported
)
