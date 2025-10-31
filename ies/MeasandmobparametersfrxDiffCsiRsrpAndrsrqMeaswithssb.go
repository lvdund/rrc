package ies

import "rrc/utils"

// MeasAndMobParametersFRX-Diff-csi-RSRP-AndRSRQ-MeasWithSSB ::= ENUMERATED
type MeasandmobparametersfrxDiffCsiRsrpAndrsrqMeaswithssb struct {
	Value utils.ENUMERATED
}

const (
	MeasandmobparametersfrxDiffCsiRsrpAndrsrqMeaswithssbEnumeratedNothing = iota
	MeasandmobparametersfrxDiffCsiRsrpAndrsrqMeaswithssbEnumeratedSupported
)
