package ies

import "rrc/utils"

// MeasResultCellSFTD-r15 ::= SEQUENCE
type MeasresultcellsftdR15 struct {
	PhyscellidR15                PhyscellidnrR15
	SfnOffsetresultR15           utils.INTEGER `lb:0,ub:1023`
	FrameboundaryoffsetresultR15 utils.INTEGER `lb:0,ub:30719`
	RsrpresultR15                *RsrpRangenrR15
}
