package ies

import "rrc/utils"

// BandCombination-v1630 ::= SEQUENCE
type BandcombinationV1630 struct {
	CaParametersnrV1630                     *CaParametersnrV1630
	CaParametersnrdcV1630                   *CaParametersnrdcV1630
	MrdcParametersV1630                     *MrdcParametersV1630
	SupportedtxbandcomblistperbcSidelinkR16 *utils.BITSTRING            `lb:1,ub:maxBandComb`
	SupportedrxbandcomblistperbcSidelinkR16 *utils.BITSTRING            `lb:1,ub:maxBandComb`
	ScalingfactortxsidelinkR16              *[]ScalingfactorsidelinkR16 `lb:1,ub:maxBandComb`
	ScalingfactorrxsidelinkR16              *[]ScalingfactorsidelinkR16 `lb:1,ub:maxBandComb`
}
