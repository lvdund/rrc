package ies

import "rrc/utils"

// MeasAndMobParametersFRX-Diff-csi-RSRP-AndRSRQ-MeasWithoutSSB ::= ENUMERATED
type MeasandmobparametersfrxDiffCsiRsrpAndrsrqMeaswithoutssb struct {
	Value utils.ENUMERATED
}

const (
	MeasandmobparametersfrxDiffCsiRsrpAndrsrqMeaswithoutssbEnumeratedNothing = iota
	MeasandmobparametersfrxDiffCsiRsrpAndrsrqMeaswithoutssbEnumeratedSupported
)
