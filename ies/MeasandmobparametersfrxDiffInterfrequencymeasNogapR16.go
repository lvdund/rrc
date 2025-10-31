package ies

import "rrc/utils"

// MeasAndMobParametersFRX-Diff-interFrequencyMeas-NoGap-r16 ::= ENUMERATED
type MeasandmobparametersfrxDiffInterfrequencymeasNogapR16 struct {
	Value utils.ENUMERATED
}

const (
	MeasandmobparametersfrxDiffInterfrequencymeasNogapR16EnumeratedNothing = iota
	MeasandmobparametersfrxDiffInterfrequencymeasNogapR16EnumeratedSupported
)
