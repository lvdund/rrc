package ies

import "rrc/utils"

// MeasAndMobParametersCommon-handoverFR2-1-FR2-2-r17 ::= ENUMERATED
type MeasandmobparameterscommonHandoverfr21Fr22R17 struct {
	Value utils.ENUMERATED
}

const (
	MeasandmobparameterscommonHandoverfr21Fr22R17EnumeratedNothing = iota
	MeasandmobparameterscommonHandoverfr21Fr22R17EnumeratedSupported
)
