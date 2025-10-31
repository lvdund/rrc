package ies

import "rrc/utils"

// CSG-ProximityIndicationParameters-r9-interFreqProximityIndication-r9 ::= ENUMERATED
type CsgProximityindicationparametersR9InterfreqproximityindicationR9 struct {
	Value utils.ENUMERATED
}

const (
	CsgProximityindicationparametersR9InterfreqproximityindicationR9EnumeratedNothing = iota
	CsgProximityindicationparametersR9InterfreqproximityindicationR9EnumeratedSupported
)
