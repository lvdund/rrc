package ies

import "rrc/utils"

// MeasAndMobParametersCommon-nr-CGI-Reporting-NEDC ::= ENUMERATED
type MeasandmobparameterscommonNrCgiReportingNedc struct {
	Value utils.ENUMERATED
}

const (
	MeasandmobparameterscommonNrCgiReportingNedcEnumeratedNothing = iota
	MeasandmobparameterscommonNrCgiReportingNedcEnumeratedSupported
)
