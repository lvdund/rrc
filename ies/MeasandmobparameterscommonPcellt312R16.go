package ies

import "rrc/utils"

// MeasAndMobParametersCommon-pcellT312-r16 ::= ENUMERATED
type MeasandmobparameterscommonPcellt312R16 struct {
	Value utils.ENUMERATED
}

const (
	MeasandmobparameterscommonPcellt312R16EnumeratedNothing = iota
	MeasandmobparameterscommonPcellt312R16EnumeratedSupported
)
