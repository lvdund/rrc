package ies

import "rrc/utils"

// SLRB-Uu-ConfigIndex-r16 ::= utils.INTEGER (1..maxNrofSLRB-r16)
type SlrbUuConfigindexR16 struct {
	Value utils.INTEGER `lb:0,ub:maxNrofSLRBR16`
}
