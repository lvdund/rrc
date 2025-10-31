package ies

import "rrc/utils"

// Q-QualMin-r9 ::= utils.INTEGER (-34..-3)
type QQualminR9 struct {
	Value utils.INTEGER `lb:0,ub:-3`
}
