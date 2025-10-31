package ies

import "rrc/utils"

// MeasAndMobParametersCommon-parallelSMTC-r17 ::= ENUMERATED
type MeasandmobparameterscommonParallelsmtcR17 struct {
	Value utils.ENUMERATED
}

const (
	MeasandmobparameterscommonParallelsmtcR17EnumeratedNothing = iota
	MeasandmobparameterscommonParallelsmtcR17EnumeratedN4
)
