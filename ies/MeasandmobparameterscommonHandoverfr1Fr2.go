package ies

import "rrc/utils"

// MeasAndMobParametersCommon-handoverFR1-FR2 ::= ENUMERATED
type MeasandmobparameterscommonHandoverfr1Fr2 struct {
	Value utils.ENUMERATED
}

const (
	MeasandmobparameterscommonHandoverfr1Fr2EnumeratedNothing = iota
	MeasandmobparameterscommonHandoverfr1Fr2EnumeratedSupported
)
