package ies

import "rrc/utils"

// MeasAndMobParametersFRX-Diff-simultaneousRxDataSSB-DiffNumerology ::= ENUMERATED
type MeasandmobparametersfrxDiffSimultaneousrxdatassbDiffnumerology struct {
	Value utils.ENUMERATED
}

const (
	MeasandmobparametersfrxDiffSimultaneousrxdatassbDiffnumerologyEnumeratedNothing = iota
	MeasandmobparametersfrxDiffSimultaneousrxdatassbDiffnumerologyEnumeratedSupported
)
