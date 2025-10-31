package ies

import "rrc/utils"

// MeasAndMobParametersCommon-concurrentMeasGapEUTRA-r17 ::= ENUMERATED
type MeasandmobparameterscommonConcurrentmeasgapeutraR17 struct {
	Value utils.ENUMERATED
}

const (
	MeasandmobparameterscommonConcurrentmeasgapeutraR17EnumeratedNothing = iota
	MeasandmobparameterscommonConcurrentmeasgapeutraR17EnumeratedSupported
)
