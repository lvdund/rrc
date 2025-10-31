package ies

import "rrc/utils"

// FeatureSetUplink-v1540-pa-PhaseDiscontinuityImpacts ::= ENUMERATED
type FeaturesetuplinkV1540PaPhasediscontinuityimpacts struct {
	Value utils.ENUMERATED
}

const (
	FeaturesetuplinkV1540PaPhasediscontinuityimpactsEnumeratedNothing = iota
	FeaturesetuplinkV1540PaPhasediscontinuityimpactsEnumeratedSupported
)
