package ies

import "rrc/utils"

// MeasAndMobParametersCommon-reportAddNeighMeasForPeriodic-r16 ::= ENUMERATED
type MeasandmobparameterscommonReportaddneighmeasforperiodicR16 struct {
	Value utils.ENUMERATED
}

const (
	MeasandmobparameterscommonReportaddneighmeasforperiodicR16EnumeratedNothing = iota
	MeasandmobparameterscommonReportaddneighmeasforperiodicR16EnumeratedSupported
)
