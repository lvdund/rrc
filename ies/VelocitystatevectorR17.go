package ies

import "rrc/utils"

// VelocityStateVector-r17 ::= utils.INTEGER (-131072..131071)
type VelocitystatevectorR17 struct {
	Value utils.INTEGER `lb:0,ub:131071`
}
