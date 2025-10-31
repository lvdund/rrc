package ies

import "rrc/utils"

// FeaturePriority-r17 ::= utils.INTEGER (0..7)
type FeaturepriorityR17 struct {
	Value utils.INTEGER `lb:0,ub:7`
}
