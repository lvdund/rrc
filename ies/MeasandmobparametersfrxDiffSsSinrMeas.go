package ies

import "rrc/utils"

// MeasAndMobParametersFRX-Diff-ss-SINR-Meas ::= ENUMERATED
type MeasandmobparametersfrxDiffSsSinrMeas struct {
	Value utils.ENUMERATED
}

const (
	MeasandmobparametersfrxDiffSsSinrMeasEnumeratedNothing = iota
	MeasandmobparametersfrxDiffSsSinrMeasEnumeratedSupported
)
