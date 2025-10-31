package ies

import "rrc/utils"

// BandNR-type2-PUSCH-RepetitionMultiSlots-v1650 ::= ENUMERATED
type BandnrType2PuschRepetitionmultislotsV1650 struct {
	Value utils.ENUMERATED
}

const (
	BandnrType2PuschRepetitionmultislotsV1650EnumeratedNothing = iota
	BandnrType2PuschRepetitionmultislotsV1650EnumeratedSupported
)
