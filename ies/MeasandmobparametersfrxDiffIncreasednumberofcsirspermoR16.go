package ies

import "rrc/utils"

// MeasAndMobParametersFRX-Diff-increasedNumberofCSIRSPerMO-r16 ::= ENUMERATED
type MeasandmobparametersfrxDiffIncreasednumberofcsirspermoR16 struct {
	Value utils.ENUMERATED
}

const (
	MeasandmobparametersfrxDiffIncreasednumberofcsirspermoR16EnumeratedNothing = iota
	MeasandmobparametersfrxDiffIncreasednumberofcsirspermoR16EnumeratedSupported
)
