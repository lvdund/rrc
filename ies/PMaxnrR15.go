package ies

import "rrc/utils"

// P-MaxNR-r15 ::= utils.INTEGER (-30..33)
type PMaxnrR15 struct {
	Value utils.INTEGER `lb:0,ub:33`
}
