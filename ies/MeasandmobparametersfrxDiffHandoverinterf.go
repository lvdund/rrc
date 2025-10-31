package ies

import "rrc/utils"

// MeasAndMobParametersFRX-Diff-handoverInterF ::= ENUMERATED
type MeasandmobparametersfrxDiffHandoverinterf struct {
	Value utils.ENUMERATED
}

const (
	MeasandmobparametersfrxDiffHandoverinterfEnumeratedNothing = iota
	MeasandmobparametersfrxDiffHandoverinterfEnumeratedSupported
)
