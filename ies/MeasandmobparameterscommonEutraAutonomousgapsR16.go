package ies

import "rrc/utils"

// MeasAndMobParametersCommon-eutra-AutonomousGaps-r16 ::= ENUMERATED
type MeasandmobparameterscommonEutraAutonomousgapsR16 struct {
	Value utils.ENUMERATED
}

const (
	MeasandmobparameterscommonEutraAutonomousgapsR16EnumeratedNothing = iota
	MeasandmobparameterscommonEutraAutonomousgapsR16EnumeratedSupported
)
