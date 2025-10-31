package ies

import "rrc/utils"

// MeasAndMobParametersMRDC-Common-independentGapConfig ::= ENUMERATED
type MeasandmobparametersmrdcCommonIndependentgapconfig struct {
	Value utils.ENUMERATED
}

const (
	MeasandmobparametersmrdcCommonIndependentgapconfigEnumeratedNothing = iota
	MeasandmobparametersmrdcCommonIndependentgapconfigEnumeratedSupported
)
