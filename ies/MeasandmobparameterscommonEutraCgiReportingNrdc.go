package ies

import "rrc/utils"

// MeasAndMobParametersCommon-eutra-CGI-Reporting-NRDC ::= ENUMERATED
type MeasandmobparameterscommonEutraCgiReportingNrdc struct {
	Value utils.ENUMERATED
}

const (
	MeasandmobparameterscommonEutraCgiReportingNrdcEnumeratedNothing = iota
	MeasandmobparameterscommonEutraCgiReportingNrdcEnumeratedSupported
)
