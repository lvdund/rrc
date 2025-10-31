package ies

import "rrc/utils"

// RejectWaitTime ::= utils.INTEGER (1..16)
type Rejectwaittime struct {
	Value utils.INTEGER `lb:0,ub:16`
}
