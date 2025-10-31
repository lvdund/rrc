package ies

import "rrc/utils"

// MeasAndMobParametersFRX-Diff-nr-AutonomousGaps-NRDC-r16 ::= ENUMERATED
type MeasandmobparametersfrxDiffNrAutonomousgapsNrdcR16 struct {
	Value utils.ENUMERATED
}

const (
	MeasandmobparametersfrxDiffNrAutonomousgapsNrdcR16EnumeratedNothing = iota
	MeasandmobparametersfrxDiffNrAutonomousgapsNrdcR16EnumeratedSupported
)
