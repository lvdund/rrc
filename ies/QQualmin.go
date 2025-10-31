package ies

import "rrc/utils"

// Q-QualMin ::= utils.INTEGER (-43..-12)
type QQualmin struct {
	Value utils.INTEGER `lb:0,ub:-12`
}
