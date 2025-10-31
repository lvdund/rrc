package ies

import "rrc/utils"

// MeasAndMobParametersCommon-parallelMeasurementGap-r17 ::= ENUMERATED
type MeasandmobparameterscommonParallelmeasurementgapR17 struct {
	Value utils.ENUMERATED
}

const (
	MeasandmobparameterscommonParallelmeasurementgapR17EnumeratedNothing = iota
	MeasandmobparameterscommonParallelmeasurementgapR17EnumeratedN2
)
