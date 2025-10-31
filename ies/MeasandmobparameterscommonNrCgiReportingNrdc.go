package ies

import "rrc/utils"

// MeasAndMobParametersCommon-nr-CGI-Reporting-NRDC ::= ENUMERATED
type MeasandmobparameterscommonNrCgiReportingNrdc struct {
	Value utils.ENUMERATED
}

const (
	MeasandmobparameterscommonNrCgiReportingNrdcEnumeratedNothing = iota
	MeasandmobparameterscommonNrCgiReportingNrdcEnumeratedSupported
)
