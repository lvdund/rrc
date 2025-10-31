package ies

import "rrc/utils"

// MeasAndMobParametersCommon-mfbi-IAB-r16 ::= ENUMERATED
type MeasandmobparameterscommonMfbiIabR16 struct {
	Value utils.ENUMERATED
}

const (
	MeasandmobparameterscommonMfbiIabR16EnumeratedNothing = iota
	MeasandmobparameterscommonMfbiIabR16EnumeratedSupported
)
