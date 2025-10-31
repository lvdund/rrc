package ies

import "rrc/utils"

// MeasAndMobParametersCommon-eutra-CGI-Reporting-NEDC ::= ENUMERATED
type MeasandmobparameterscommonEutraCgiReportingNedc struct {
	Value utils.ENUMERATED
}

const (
	MeasandmobparameterscommonEutraCgiReportingNedcEnumeratedNothing = iota
	MeasandmobparameterscommonEutraCgiReportingNedcEnumeratedSupported
)
