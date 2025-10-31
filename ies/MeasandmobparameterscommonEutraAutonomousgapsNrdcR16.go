package ies

import "rrc/utils"

// MeasAndMobParametersCommon-eutra-AutonomousGaps-NRDC-r16 ::= ENUMERATED
type MeasandmobparameterscommonEutraAutonomousgapsNrdcR16 struct {
	Value utils.ENUMERATED
}

const (
	MeasandmobparameterscommonEutraAutonomousgapsNrdcR16EnumeratedNothing = iota
	MeasandmobparameterscommonEutraAutonomousgapsNrdcR16EnumeratedSupported
)
