package ies

import "rrc/utils"

// MeasAndMobParametersCommon-gNB-ID-LengthReporting-NPN-r17 ::= ENUMERATED
type MeasandmobparameterscommonGnbIdLengthreportingNpnR17 struct {
	Value utils.ENUMERATED
}

const (
	MeasandmobparameterscommonGnbIdLengthreportingNpnR17EnumeratedNothing = iota
	MeasandmobparameterscommonGnbIdLengthreportingNpnR17EnumeratedSupported
)
