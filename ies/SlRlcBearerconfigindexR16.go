package ies

import "rrc/utils"

// SL-RLC-BearerConfigIndex-r16 ::= utils.INTEGER (1..maxSL-LCID-r16)
type SlRlcBearerconfigindexR16 struct {
	Value utils.INTEGER `lb:0,ub:maxSLLcidR16`
}
