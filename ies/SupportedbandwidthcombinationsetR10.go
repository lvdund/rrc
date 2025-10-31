package ies

import "rrc/utils"

// SupportedBandwidthCombinationSet-r10 ::= BIT STRING (SIZE (1..maxBandwidthCombSet-r10))
type SupportedbandwidthcombinationsetR10 struct {
	Value utils.BITSTRING `lb:1,ub:maxBandwidthCombSetR10`
}
