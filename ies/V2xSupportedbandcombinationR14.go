package ies

import "rrc/utils"

// V2X-SupportedBandCombination-r14 ::= SEQUENCE OF V2X-BandCombinationParameters-r14
// SIZE (1..maxBandComb-r13)
type V2xSupportedbandcombinationR14 struct {
	Value []V2xBandcombinationparametersR14
}
