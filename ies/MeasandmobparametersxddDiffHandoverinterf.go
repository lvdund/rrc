package ies

import "rrc/utils"

// MeasAndMobParametersXDD-Diff-handoverInterF ::= ENUMERATED
type MeasandmobparametersxddDiffHandoverinterf struct {
	Value utils.ENUMERATED
}

const (
	MeasandmobparametersxddDiffHandoverinterfEnumeratedNothing = iota
	MeasandmobparametersxddDiffHandoverinterfEnumeratedSupported
)
