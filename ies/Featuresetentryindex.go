package ies

import "rrc/utils"

// FeatureSetEntryIndex ::= utils.INTEGER (1.. maxFeatureSetsPerBand)
type Featuresetentryindex struct {
	Value utils.INTEGER `lb:0,ub:maxFeatureSetsPerBand`
}
