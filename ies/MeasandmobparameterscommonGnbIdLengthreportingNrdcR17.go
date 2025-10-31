package ies

import "rrc/utils"

// MeasAndMobParametersCommon-gNB-ID-LengthReporting-NRDC-r17 ::= ENUMERATED
type MeasandmobparameterscommonGnbIdLengthreportingNrdcR17 struct {
	Value utils.ENUMERATED
}

const (
	MeasandmobparameterscommonGnbIdLengthreportingNrdcR17EnumeratedNothing = iota
	MeasandmobparameterscommonGnbIdLengthreportingNrdcR17EnumeratedSupported
)
