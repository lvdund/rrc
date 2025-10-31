package ies

import "rrc/utils"

// MeasAndMobParametersMRDC-v1610-interNR-MeasEUTRA-IAB-r16 ::= ENUMERATED
type MeasandmobparametersmrdcV1610InternrMeaseutraIabR16 struct {
	Value utils.ENUMERATED
}

const (
	MeasandmobparametersmrdcV1610InternrMeaseutraIabR16EnumeratedNothing = iota
	MeasandmobparametersmrdcV1610InternrMeaseutraIabR16EnumeratedSupported
)
