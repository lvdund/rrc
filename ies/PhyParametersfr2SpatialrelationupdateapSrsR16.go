package ies

import "rrc/utils"

// Phy-ParametersFR2-spatialRelationUpdateAP-SRS-r16 ::= ENUMERATED
type PhyParametersfr2SpatialrelationupdateapSrsR16 struct {
	Value utils.ENUMERATED
}

const (
	PhyParametersfr2SpatialrelationupdateapSrsR16EnumeratedNothing = iota
	PhyParametersfr2SpatialrelationupdateapSrsR16EnumeratedSupported
)
