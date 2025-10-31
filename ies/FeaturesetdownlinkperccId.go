package ies

import "rrc/utils"

// FeatureSetDownlinkPerCC-Id ::= utils.INTEGER (1..maxPerCC-FeatureSets)
type FeaturesetdownlinkperccId struct {
	Value utils.INTEGER `lb:0,ub:maxPerCCFeaturesets`
}
