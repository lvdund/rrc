package ies

import "rrc/utils"

// PRACH-ParametersCE-r13 ::= SEQUENCE
type PrachParametersceR13 struct {
	PrachConfigindexR13                utils.INTEGER `lb:0,ub:63`
	PrachFreqoffsetR13                 utils.INTEGER `lb:0,ub:94`
	PrachStartingsubframeR13           *PrachParametersceR13PrachStartingsubframeR13
	MaxnumpreambleattemptceR13         *PrachParametersceR13MaxnumpreambleattemptceR13
	NumrepetitionperpreambleattemptR13 PrachParametersceR13NumrepetitionperpreambleattemptR13
	MpdcchNarrowbandstomonitorR13      []utils.INTEGER `lb:1,ub:2`
	MpdcchNumrepetitionRaR13           PrachParametersceR13MpdcchNumrepetitionRaR13
	PrachHoppingconfigR13              PrachParametersceR13PrachHoppingconfigR13
}
