package ies

import "rrc/utils"

// BandCombinationParameters-v1250 ::= SEQUENCE
// Extensible
type BandcombinationparametersV1250 struct {
	DcSupportR12               *BandcombinationparametersV1250DcSupportR12
	Supportednaics2crsApR12    *utils.BITSTRING `lb:1,ub:maxNAICSEntriesR12`
	CommsupportedbandsperbcR12 *utils.BITSTRING `lb:1,ub:maxBands`
}
