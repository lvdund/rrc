package ies

import "rrc/utils"

// MeasResultSSTD-r13 ::= SEQUENCE
type MeasresultsstdR13 struct {
	SfnOffsetresultR13              utils.INTEGER
	FrameboundaryoffsetresultR13    utils.INTEGER
	SubframeboundaryoffsetresultR13 utils.INTEGER
}
