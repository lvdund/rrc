package ies

import "rrc/utils"

// BandNR-pusch-RepetitionMultiSlots-v1650 ::= ENUMERATED
type BandnrPuschRepetitionmultislotsV1650 struct {
	Value utils.ENUMERATED
}

const (
	BandnrPuschRepetitionmultislotsV1650EnumeratedNothing = iota
	BandnrPuschRepetitionmultislotsV1650EnumeratedSupported
)
