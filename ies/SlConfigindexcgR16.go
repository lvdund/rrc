package ies

import "rrc/utils"

// SL-ConfigIndexCG-r16 ::= utils.INTEGER (0..maxNrofCG-SL-1-r16)
type SlConfigindexcgR16 struct {
	Value utils.INTEGER `lb:0,ub:maxNrofCGSl1R16`
}
