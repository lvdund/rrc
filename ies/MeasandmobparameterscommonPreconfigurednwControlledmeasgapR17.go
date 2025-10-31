package ies

import "rrc/utils"

// MeasAndMobParametersCommon-preconfiguredNW-ControlledMeasGap-r17 ::= ENUMERATED
type MeasandmobparameterscommonPreconfigurednwControlledmeasgapR17 struct {
	Value utils.ENUMERATED
}

const (
	MeasandmobparameterscommonPreconfigurednwControlledmeasgapR17EnumeratedNothing = iota
	MeasandmobparameterscommonPreconfigurednwControlledmeasgapR17EnumeratedSupported
)
