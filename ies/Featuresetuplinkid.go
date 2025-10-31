package ies

import "rrc/utils"

// FeatureSetUplinkId ::= utils.INTEGER (0..maxUplinkFeatureSets)
type Featuresetuplinkid struct {
	Value utils.INTEGER `lb:0,ub:maxUplinkFeatureSets`
}
