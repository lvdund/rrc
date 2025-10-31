package ies

import "rrc/utils"

// MeasAndMobParametersFRX-Diff-cli-RSSI-Meas-r16 ::= ENUMERATED
type MeasandmobparametersfrxDiffCliRssiMeasR16 struct {
	Value utils.ENUMERATED
}

const (
	MeasandmobparametersfrxDiffCliRssiMeasR16EnumeratedNothing = iota
	MeasandmobparametersfrxDiffCliRssiMeasR16EnumeratedSupported
)
