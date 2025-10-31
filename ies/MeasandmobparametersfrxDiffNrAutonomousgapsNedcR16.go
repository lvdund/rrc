package ies

import "rrc/utils"

// MeasAndMobParametersFRX-Diff-nr-AutonomousGaps-NEDC-r16 ::= ENUMERATED
type MeasandmobparametersfrxDiffNrAutonomousgapsNedcR16 struct {
	Value utils.ENUMERATED
}

const (
	MeasandmobparametersfrxDiffNrAutonomousgapsNedcR16EnumeratedNothing = iota
	MeasandmobparametersfrxDiffNrAutonomousgapsNedcR16EnumeratedSupported
)
