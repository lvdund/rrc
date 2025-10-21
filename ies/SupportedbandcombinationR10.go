package ies

import "rrc/utils"

// SupportedBandCombination-r10 ::= SEQUENCE OF BandCombinationParameters-r10
// SIZE (1..maxBandComb-r10)
type SupportedbandcombinationR10 struct {
	Value []BandcombinationparametersR10
}
