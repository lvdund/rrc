package ies

import "rrc/utils"

// MeasAndMobParametersFRX-Diff-idleInactiveNR-MeasReport-r16 ::= ENUMERATED
type MeasandmobparametersfrxDiffIdleinactivenrMeasreportR16 struct {
	Value utils.ENUMERATED
}

const (
	MeasandmobparametersfrxDiffIdleinactivenrMeasreportR16EnumeratedNothing = iota
	MeasandmobparametersfrxDiffIdleinactivenrMeasreportR16EnumeratedSupported
)
