package ies

import "rrc/utils"

// MeasAndMobParametersCommon-supportedGapPattern-NRonly-NEDC-r16 ::= ENUMERATED
type MeasandmobparameterscommonSupportedgappatternNronlyNedcR16 struct {
	Value utils.ENUMERATED
}

const (
	MeasandmobparameterscommonSupportedgappatternNronlyNedcR16EnumeratedNothing = iota
	MeasandmobparameterscommonSupportedgappatternNronlyNedcR16EnumeratedSupported
)
