package ies

import "rrc/utils"

// MeasAndMobParametersCommon-nr-CGI-Reporting-ENDC ::= ENUMERATED
type MeasandmobparameterscommonNrCgiReportingEndc struct {
	Value utils.ENUMERATED
}

const (
	MeasandmobparameterscommonNrCgiReportingEndcEnumeratedNothing = iota
	MeasandmobparameterscommonNrCgiReportingEndcEnumeratedSupported
)
