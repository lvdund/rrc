package ies

import "rrc/utils"

// MeasResultSFTD-EUTRA ::= SEQUENCE
type MeasresultsftdEutra struct {
	EutraPhyscellid           EutraPhyscellid
	SfnOffsetresult           utils.INTEGER `lb:0,ub:1023`
	Frameboundaryoffsetresult utils.INTEGER `lb:0,ub:30719`
	RsrpResult                *RsrpRange
}
