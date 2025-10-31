package ies

import "rrc/utils"

// BandCombination-v1590 ::= SEQUENCE
type BandcombinationV1590 struct {
	Supportedbandwidthcombinationsetintraendc *utils.BITSTRING `lb:1,ub:32`
	MrdcParametersV1590                       MrdcParametersV1590
}
