package ies

import "rrc/utils"

// MeasAndMobParametersCommon-preconfiguredUE-AutonomousMeasGap-r17 ::= ENUMERATED
type MeasandmobparameterscommonPreconfiguredueAutonomousmeasgapR17 struct {
	Value utils.ENUMERATED
}

const (
	MeasandmobparameterscommonPreconfiguredueAutonomousmeasgapR17EnumeratedNothing = iota
	MeasandmobparameterscommonPreconfiguredueAutonomousmeasgapR17EnumeratedSupported
)
