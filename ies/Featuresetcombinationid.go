package ies

import "rrc/utils"

// FeatureSetCombinationId ::= utils.INTEGER (0.. maxFeatureSetCombinations)
type Featuresetcombinationid struct {
	Value utils.INTEGER `lb:0,ub:maxFeatureSetCombinations`
}
