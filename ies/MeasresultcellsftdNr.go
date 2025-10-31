package ies

import "rrc/utils"

// MeasResultCellSFTD-NR ::= SEQUENCE
type MeasresultcellsftdNr struct {
	Physcellid                Physcellid
	SfnOffsetresult           utils.INTEGER `lb:0,ub:1023`
	Frameboundaryoffsetresult utils.INTEGER `lb:0,ub:30719`
	RsrpResult                *RsrpRange
}
