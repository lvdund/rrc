package ies

import "rrc/utils"

// MeasAndMobParametersCommon-nr-CGI-Reporting-NPN-r16 ::= ENUMERATED
type MeasandmobparameterscommonNrCgiReportingNpnR16 struct {
	Value utils.ENUMERATED
}

const (
	MeasandmobparameterscommonNrCgiReportingNpnR16EnumeratedNothing = iota
	MeasandmobparameterscommonNrCgiReportingNpnR16EnumeratedSupported
)
