package ies

import "rrc/utils"

// MeasAndMobParametersCommon-ncsg-MeasGapPerFR-r17 ::= ENUMERATED
type MeasandmobparameterscommonNcsgMeasgapperfrR17 struct {
	Value utils.ENUMERATED
}

const (
	MeasandmobparameterscommonNcsgMeasgapperfrR17EnumeratedNothing = iota
	MeasandmobparameterscommonNcsgMeasgapperfrR17EnumeratedSupported
)
