package ies

import "rrc/utils"

// MeasAndMobParametersCommon-condHandoverParametersCommon-r16-condHandoverFR1-FR2-r16 ::= ENUMERATED
type MeasandmobparameterscommonCondhandoverparameterscommonR16Condhandoverfr1Fr2R16 struct {
	Value utils.ENUMERATED
}

const (
	MeasandmobparameterscommonCondhandoverparameterscommonR16Condhandoverfr1Fr2R16EnumeratedNothing = iota
	MeasandmobparameterscommonCondhandoverparameterscommonR16Condhandoverfr1Fr2R16EnumeratedSupported
)
