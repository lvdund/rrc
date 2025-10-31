package ies

import "rrc/utils"

// ReducedMaxCCs-r16 ::= SEQUENCE
type ReducedmaxccsR16 struct {
	ReducedccsdlR16 utils.INTEGER `lb:0,ub:31`
	ReducedccsulR16 utils.INTEGER `lb:0,ub:31`
}
