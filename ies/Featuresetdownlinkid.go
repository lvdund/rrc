package ies

import "rrc/utils"

// FeatureSetDownlinkId ::= utils.INTEGER (0..maxDownlinkFeatureSets)
type Featuresetdownlinkid struct {
	Value utils.INTEGER `lb:0,ub:maxDownlinkFeatureSets`
}
