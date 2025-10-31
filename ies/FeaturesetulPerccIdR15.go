package ies

import "rrc/utils"

// FeatureSetUL-PerCC-Id-r15 ::= utils.INTEGER (0..maxPerCC-FeatureSets-r15)
type FeaturesetulPerccIdR15 struct {
	Value utils.INTEGER `lb:0,ub:maxPerCCFeaturesetsR15`
}
