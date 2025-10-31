package ies

import "rrc/utils"

// CSG-ProximityIndicationParameters-r9-intraFreqProximityIndication-r9 ::= ENUMERATED
type CsgProximityindicationparametersR9IntrafreqproximityindicationR9 struct {
	Value utils.ENUMERATED
}

const (
	CsgProximityindicationparametersR9IntrafreqproximityindicationR9EnumeratedNothing = iota
	CsgProximityindicationparametersR9IntrafreqproximityindicationR9EnumeratedSupported
)
