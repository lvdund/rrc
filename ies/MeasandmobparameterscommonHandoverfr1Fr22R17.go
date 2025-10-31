package ies

import "rrc/utils"

// MeasAndMobParametersCommon-handoverFR1-FR2-2-r17 ::= ENUMERATED
type MeasandmobparameterscommonHandoverfr1Fr22R17 struct {
	Value utils.ENUMERATED
}

const (
	MeasandmobparameterscommonHandoverfr1Fr22R17EnumeratedNothing = iota
	MeasandmobparameterscommonHandoverfr1Fr22R17EnumeratedSupported
)
