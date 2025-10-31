package ies

import "rrc/utils"

// MeasAndMobParametersCommon-eutra-CGI-Reporting ::= ENUMERATED
type MeasandmobparameterscommonEutraCgiReporting struct {
	Value utils.ENUMERATED
}

const (
	MeasandmobparameterscommonEutraCgiReportingEnumeratedNothing = iota
	MeasandmobparameterscommonEutraCgiReportingEnumeratedSupported
)
