package ies

// MeasAndMobParametersMRDC ::= SEQUENCE
type Measandmobparametersmrdc struct {
	MeasandmobparametersmrdcCommon  *MeasandmobparametersmrdcCommon
	MeasandmobparametersmrdcXddDiff *MeasandmobparametersmrdcXddDiff
	MeasandmobparametersmrdcFrxDiff *MeasandmobparametersmrdcFrxDiff
}
