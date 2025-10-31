package ies

import "rrc/utils"

// FeatureSetUplinkPerCC-Id ::= utils.INTEGER (1..maxPerCC-FeatureSets)
type FeaturesetuplinkperccId struct {
	Value utils.INTEGER `lb:0,ub:maxPerCCFeaturesets`
}
