package ies

import "rrc/utils"

// MeasAndMobParametersCommon-gNB-ID-LengthReporting-NEDC-r17 ::= ENUMERATED
type MeasandmobparameterscommonGnbIdLengthreportingNedcR17 struct {
	Value utils.ENUMERATED
}

const (
	MeasandmobparameterscommonGnbIdLengthreportingNedcR17EnumeratedNothing = iota
	MeasandmobparameterscommonGnbIdLengthreportingNedcR17EnumeratedSupported
)
