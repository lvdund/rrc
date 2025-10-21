package ies

import "rrc/utils"

// BandCombinationParameters-v1250 ::= SEQUENCE
// Extensible
type BandcombinationparametersV1250 struct {
	DcSupportR12               *interface{}
	Supportednaics2crsApR12    *utils.BITSTRING
	CommsupportedbandsperbcR12 *utils.BITSTRING
}
