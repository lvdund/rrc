package ies

import "rrc/utils"

// MeasAndMobParametersMRDC-Common-v1610-pscellT312-r16 ::= ENUMERATED
type MeasandmobparametersmrdcCommonV1610Pscellt312R16 struct {
	Value utils.ENUMERATED
}

const (
	MeasandmobparametersmrdcCommonV1610Pscellt312R16EnumeratedNothing = iota
	MeasandmobparametersmrdcCommonV1610Pscellt312R16EnumeratedSupported
)
