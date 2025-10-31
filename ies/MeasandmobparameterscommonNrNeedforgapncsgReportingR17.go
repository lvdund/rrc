package ies

import "rrc/utils"

// MeasAndMobParametersCommon-nr-NeedForGapNCSG-Reporting-r17 ::= ENUMERATED
type MeasandmobparameterscommonNrNeedforgapncsgReportingR17 struct {
	Value utils.ENUMERATED
}

const (
	MeasandmobparameterscommonNrNeedforgapncsgReportingR17EnumeratedNothing = iota
	MeasandmobparameterscommonNrNeedforgapncsgReportingR17EnumeratedSupported
)
