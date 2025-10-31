package ies

import "rrc/utils"

// EDT-PRACH-ParametersCE-r15-edt-PRACH-ParametersCE-r15 ::= SEQUENCE
type EdtPrachParametersceR15EdtPrachParametersceR15 struct {
	PrachConfigindexR15           utils.INTEGER `lb:0,ub:63`
	PrachFreqoffsetR15            utils.INTEGER `lb:0,ub:94`
	PrachStartingsubframeR15      *EdtPrachParametersceR15EdtPrachParametersceR15PrachStartingsubframeR15
	MpdcchNarrowbandstomonitorR15 []utils.INTEGER `lb:1,ub:2`
}
