package ies

import "rrc/utils"

// MeasAndMobParametersCommon-eutra-NeedForGapNCSG-Reporting-r17 ::= ENUMERATED
type MeasandmobparameterscommonEutraNeedforgapncsgReportingR17 struct {
	Value utils.ENUMERATED
}

const (
	MeasandmobparameterscommonEutraNeedforgapncsgReportingR17EnumeratedNothing = iota
	MeasandmobparameterscommonEutraNeedforgapncsgReportingR17EnumeratedSupported
)
