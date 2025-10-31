package ies

import "rrc/utils"

// SCS-SpecificDuration-r17 ::= utils.INTEGER (1..166)
type ScsSpecificdurationR17 struct {
	Value utils.INTEGER `lb:0,ub:166`
}
