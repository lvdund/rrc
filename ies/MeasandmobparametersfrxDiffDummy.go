package ies

import "rrc/utils"

// MeasAndMobParametersFRX-Diff-dummy ::= ENUMERATED
type MeasandmobparametersfrxDiffDummy struct {
	Value utils.ENUMERATED
}

const (
	MeasandmobparametersfrxDiffDummyEnumeratedNothing = iota
	MeasandmobparametersfrxDiffDummyEnumeratedSupported
)
