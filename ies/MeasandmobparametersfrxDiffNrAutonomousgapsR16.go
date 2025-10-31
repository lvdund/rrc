package ies

import "rrc/utils"

// MeasAndMobParametersFRX-Diff-nr-AutonomousGaps-r16 ::= ENUMERATED
type MeasandmobparametersfrxDiffNrAutonomousgapsR16 struct {
	Value utils.ENUMERATED
}

const (
	MeasandmobparametersfrxDiffNrAutonomousgapsR16EnumeratedNothing = iota
	MeasandmobparametersfrxDiffNrAutonomousgapsR16EnumeratedSupported
)
