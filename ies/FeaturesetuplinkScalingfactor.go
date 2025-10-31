package ies

import "rrc/utils"

// FeatureSetUplink-scalingFactor ::= ENUMERATED
type FeaturesetuplinkScalingfactor struct {
	Value utils.ENUMERATED
}

const (
	FeaturesetuplinkScalingfactorEnumeratedNothing = iota
	FeaturesetuplinkScalingfactorEnumeratedF0p4
	FeaturesetuplinkScalingfactorEnumeratedF0p75
	FeaturesetuplinkScalingfactorEnumeratedF0p8
)
