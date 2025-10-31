package ies

import "rrc/utils"

// TimeConnSourceDAPS-Failure-r17 ::= utils.INTEGER (0..1023)
type TimeconnsourcedapsFailureR17 struct {
	Value utils.INTEGER `lb:0,ub:1023`
}
