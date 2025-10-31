package ies

import "rrc/utils"

// MeasAndMobParametersCommon-nr-NeedForGap-Reporting-r16 ::= ENUMERATED
type MeasandmobparameterscommonNrNeedforgapReportingR16 struct {
	Value utils.ENUMERATED
}

const (
	MeasandmobparameterscommonNrNeedforgapReportingR16EnumeratedNothing = iota
	MeasandmobparameterscommonNrNeedforgapReportingR16EnumeratedSupported
)
