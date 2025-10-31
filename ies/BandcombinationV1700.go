package ies

import "rrc/utils"

// BandCombination-v1700 ::= SEQUENCE
type BandcombinationV1700 struct {
	CaParametersnrV1700                              *CaParametersnrV1700
	CaParametersnrdcV1700                            *CaParametersnrdcV1700
	MrdcParametersV1700                              *MrdcParametersV1700
	BandlistV1710                                    *[]BandparametersV1710 `lb:1,ub:maxSimultaneousBands`
	SupportedbandcomblistperbcSlRelaydiscoveryR17    *utils.BITSTRING       `lb:1,ub:maxBandComb`
	SupportedbandcomblistperbcSlNonrelaydiscoveryR17 *utils.BITSTRING       `lb:1,ub:maxBandComb`
}
