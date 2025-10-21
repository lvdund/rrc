package ies

import "rrc/utils"

// BandCombinationParameters-v1430 ::= SEQUENCE
type BandcombinationparametersV1430 struct {
	BandparameterlistV1430             *interface{}
	V2xSupportedtxbandcomblistperbcR14 *utils.BITSTRING
	V2xSupportedrxbandcomblistperbcR14 *utils.BITSTRING
}
