package ies

import "rrc/utils"

// MeasAndMobParametersFRX-Diff-idleInactiveNR-MeasBeamReport-r16 ::= ENUMERATED
type MeasandmobparametersfrxDiffIdleinactivenrMeasbeamreportR16 struct {
	Value utils.ENUMERATED
}

const (
	MeasandmobparametersfrxDiffIdleinactivenrMeasbeamreportR16EnumeratedNothing = iota
	MeasandmobparametersfrxDiffIdleinactivenrMeasbeamreportR16EnumeratedSupported
)
