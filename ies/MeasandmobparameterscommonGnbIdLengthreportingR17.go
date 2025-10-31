package ies

import "rrc/utils"

// MeasAndMobParametersCommon-gNB-ID-LengthReporting-r17 ::= ENUMERATED
type MeasandmobparameterscommonGnbIdLengthreportingR17 struct {
	Value utils.ENUMERATED
}

const (
	MeasandmobparameterscommonGnbIdLengthreportingR17EnumeratedNothing = iota
	MeasandmobparameterscommonGnbIdLengthreportingR17EnumeratedSupported
)
