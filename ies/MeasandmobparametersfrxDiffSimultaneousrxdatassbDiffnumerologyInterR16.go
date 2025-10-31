package ies

import "rrc/utils"

// MeasAndMobParametersFRX-Diff-simultaneousRxDataSSB-DiffNumerology-Inter-r16 ::= ENUMERATED
type MeasandmobparametersfrxDiffSimultaneousrxdatassbDiffnumerologyInterR16 struct {
	Value utils.ENUMERATED
}

const (
	MeasandmobparametersfrxDiffSimultaneousrxdatassbDiffnumerologyInterR16EnumeratedNothing = iota
	MeasandmobparametersfrxDiffSimultaneousrxdatassbDiffnumerologyInterR16EnumeratedSupported
)
