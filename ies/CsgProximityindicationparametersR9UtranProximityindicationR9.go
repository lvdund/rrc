package ies

import "rrc/utils"

// CSG-ProximityIndicationParameters-r9-utran-ProximityIndication-r9 ::= ENUMERATED
type CsgProximityindicationparametersR9UtranProximityindicationR9 struct {
	Value utils.ENUMERATED
}

const (
	CsgProximityindicationparametersR9UtranProximityindicationR9EnumeratedNothing = iota
	CsgProximityindicationparametersR9UtranProximityindicationR9EnumeratedSupported
)
