package ies

import "rrc/utils"

// MeasAndMobParametersMRDC-FRX-Diff-simultaneousRxDataSSB-DiffNumerology ::= ENUMERATED
type MeasandmobparametersmrdcFrxDiffSimultaneousrxdatassbDiffnumerology struct {
	Value utils.ENUMERATED
}

const (
	MeasandmobparametersmrdcFrxDiffSimultaneousrxdatassbDiffnumerologyEnumeratedNothing = iota
	MeasandmobparametersmrdcFrxDiffSimultaneousrxdatassbDiffnumerologyEnumeratedSupported
)
