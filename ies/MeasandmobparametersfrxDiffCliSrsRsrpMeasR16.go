package ies

import "rrc/utils"

// MeasAndMobParametersFRX-Diff-cli-SRS-RSRP-Meas-r16 ::= ENUMERATED
type MeasandmobparametersfrxDiffCliSrsRsrpMeasR16 struct {
	Value utils.ENUMERATED
}

const (
	MeasandmobparametersfrxDiffCliSrsRsrpMeasR16EnumeratedNothing = iota
	MeasandmobparametersfrxDiffCliSrsRsrpMeasR16EnumeratedSupported
)
