package ies

import "rrc/utils"

// MeasAndMobParametersCommon-nr-CGI-Reporting ::= ENUMERATED
type MeasandmobparameterscommonNrCgiReporting struct {
	Value utils.ENUMERATED
}

const (
	MeasandmobparameterscommonNrCgiReportingEnumeratedNothing = iota
	MeasandmobparameterscommonNrCgiReportingEnumeratedSupported
)
