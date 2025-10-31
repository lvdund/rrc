package ies

import "rrc/utils"

// MeasAndMobParametersXDD-Diff-dummy ::= ENUMERATED
type MeasandmobparametersxddDiffDummy struct {
	Value utils.ENUMERATED
}

const (
	MeasandmobparametersxddDiffDummyEnumeratedNothing = iota
	MeasandmobparametersxddDiffDummyEnumeratedSupported
)
