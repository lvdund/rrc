package ies

import "rrc/utils"

// MeasAndMobParametersFRX-Diff-csi-SINR-Meas ::= ENUMERATED
type MeasandmobparametersfrxDiffCsiSinrMeas struct {
	Value utils.ENUMERATED
}

const (
	MeasandmobparametersfrxDiffCsiSinrMeasEnumeratedNothing = iota
	MeasandmobparametersfrxDiffCsiSinrMeasEnumeratedSupported
)
