package ies

import "rrc/utils"

// MeasAndMobParametersCommon-gNB-ID-LengthReporting-ENDC-r17 ::= ENUMERATED
type MeasandmobparameterscommonGnbIdLengthreportingEndcR17 struct {
	Value utils.ENUMERATED
}

const (
	MeasandmobparameterscommonGnbIdLengthreportingEndcR17EnumeratedNothing = iota
	MeasandmobparameterscommonGnbIdLengthreportingEndcR17EnumeratedSupported
)
