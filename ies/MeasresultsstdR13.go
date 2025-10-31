package ies

import "rrc/utils"

// MeasResultSSTD-r13 ::= SEQUENCE
type MeasresultsstdR13 struct {
	SfnOffsetresultR13              utils.INTEGER `lb:0,ub:1023`
	FrameboundaryoffsetresultR13    utils.INTEGER `lb:0,ub:4`
	SubframeboundaryoffsetresultR13 utils.INTEGER `lb:0,ub:127`
}
