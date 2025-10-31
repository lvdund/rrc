package ies

import "rrc/utils"

// MIMO-ParametersPerBand-supportCodeWordSoftCombining-r16 ::= ENUMERATED
type MimoParametersperbandSupportcodewordsoftcombiningR16 struct {
	Value utils.ENUMERATED
}

const (
	MimoParametersperbandSupportcodewordsoftcombiningR16EnumeratedNothing = iota
	MimoParametersperbandSupportcodewordsoftcombiningR16EnumeratedSupported
)
