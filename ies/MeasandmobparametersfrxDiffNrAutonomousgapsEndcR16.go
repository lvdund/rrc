package ies

import "rrc/utils"

// MeasAndMobParametersFRX-Diff-nr-AutonomousGaps-ENDC-r16 ::= ENUMERATED
type MeasandmobparametersfrxDiffNrAutonomousgapsEndcR16 struct {
	Value utils.ENUMERATED
}

const (
	MeasandmobparametersfrxDiffNrAutonomousgapsEndcR16EnumeratedNothing = iota
	MeasandmobparametersfrxDiffNrAutonomousgapsEndcR16EnumeratedSupported
)
