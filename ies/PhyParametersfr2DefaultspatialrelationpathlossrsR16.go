package ies

import "rrc/utils"

// Phy-ParametersFR2-defaultSpatialRelationPathlossRS-r16 ::= ENUMERATED
type PhyParametersfr2DefaultspatialrelationpathlossrsR16 struct {
	Value utils.ENUMERATED
}

const (
	PhyParametersfr2DefaultspatialrelationpathlossrsR16EnumeratedNothing = iota
	PhyParametersfr2DefaultspatialrelationpathlossrsR16EnumeratedSupported
)
