package ies

import "rrc/utils"

// MeasAndMobParametersCommon-dummy ::= ENUMERATED
type MeasandmobparameterscommonDummy struct {
	Value utils.ENUMERATED
}

const (
	MeasandmobparameterscommonDummyEnumeratedNothing = iota
	MeasandmobparameterscommonDummyEnumeratedSupported
)
