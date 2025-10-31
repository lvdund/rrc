package ies

import "rrc/utils"

// MeasAndMobParametersMRDC-Common-v1730-independentGapConfig-maxCC-r17 ::= SEQUENCE
type MeasandmobparametersmrdcCommonV1730IndependentgapconfigMaxccR17 struct {
	Fr1OnlyR17   *utils.INTEGER `lb:0,ub:32`
	Fr2OnlyR17   *utils.INTEGER `lb:0,ub:32`
	Fr1Andfr2R17 *utils.INTEGER `lb:0,ub:32`
}
