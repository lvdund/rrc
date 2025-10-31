package ies

import "rrc/utils"

// SL-Freq-Id-r16 ::= utils.INTEGER (1.. maxNrofFreqSL-r16)
type SlFreqIdR16 struct {
	Value utils.INTEGER `lb:0,ub:maxNrofFreqSLR16`
}
