package ies

import "rrc/utils"

// MeasAndMobParametersCommon-eutra-AutonomousGaps-NEDC-r16 ::= ENUMERATED
type MeasandmobparameterscommonEutraAutonomousgapsNedcR16 struct {
	Value utils.ENUMERATED
}

const (
	MeasandmobparameterscommonEutraAutonomousgapsNedcR16EnumeratedNothing = iota
	MeasandmobparameterscommonEutraAutonomousgapsNedcR16EnumeratedSupported
)
