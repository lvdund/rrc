package ies

import "rrc/utils"

// BandCombinationParameters-v1430 ::= SEQUENCE
type BandcombinationparametersV1430 struct {
	BandparameterlistV1430             *[]BandparametersV1430 `lb:1,ub:maxSimultaneousBandsR10`
	V2xSupportedtxbandcomblistperbcR14 *utils.BITSTRING       `lb:1,ub:maxBandCombR13`
	V2xSupportedrxbandcomblistperbcR14 *utils.BITSTRING       `lb:1,ub:maxBandCombR13`
}
