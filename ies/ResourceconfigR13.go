package ies

import "rrc/utils"

// ResourceConfig-r13 ::= utils.INTEGER (0..31)
type ResourceconfigR13 struct {
	Value utils.INTEGER `lb:0,ub:31`
}
