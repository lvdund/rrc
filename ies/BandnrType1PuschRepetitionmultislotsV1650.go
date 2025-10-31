package ies

import "rrc/utils"

// BandNR-type1-PUSCH-RepetitionMultiSlots-v1650 ::= ENUMERATED
type BandnrType1PuschRepetitionmultislotsV1650 struct {
	Value utils.ENUMERATED
}

const (
	BandnrType1PuschRepetitionmultislotsV1650EnumeratedNothing = iota
	BandnrType1PuschRepetitionmultislotsV1650EnumeratedSupported
)
