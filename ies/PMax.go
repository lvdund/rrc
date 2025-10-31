package ies

import "rrc/utils"

// P-Max ::= utils.INTEGER (-30..33)
type PMax struct {
	Value utils.INTEGER `lb:0,ub:33`
}
