package ies

import "rrc/utils"

// MeasAndMobParametersCommon-independentGapConfig ::= ENUMERATED
type MeasandmobparameterscommonIndependentgapconfig struct {
	Value utils.ENUMERATED
}

const (
	MeasandmobparameterscommonIndependentgapconfigEnumeratedNothing = iota
	MeasandmobparameterscommonIndependentgapconfigEnumeratedSupported
)
