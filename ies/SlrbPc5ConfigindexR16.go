package ies

import "rrc/utils"

// SLRB-PC5-ConfigIndex-r16 ::= utils.INTEGER (1..maxNrofSLRB-r16)
type SlrbPc5ConfigindexR16 struct {
	Value utils.INTEGER `lb:0,ub:maxNrofSLRBR16`
}
