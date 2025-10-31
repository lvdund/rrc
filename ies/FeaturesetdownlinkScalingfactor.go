package ies

import "rrc/utils"

// FeatureSetDownlink-scalingFactor ::= ENUMERATED
type FeaturesetdownlinkScalingfactor struct {
	Value utils.ENUMERATED
}

const (
	FeaturesetdownlinkScalingfactorEnumeratedNothing = iota
	FeaturesetdownlinkScalingfactorEnumeratedF0p4
	FeaturesetdownlinkScalingfactorEnumeratedF0p75
	FeaturesetdownlinkScalingfactorEnumeratedF0p8
)
