package ies

import "rrc/utils"

// BandNR-pusch-RepetitionMsg3-r17 ::= ENUMERATED
type BandnrPuschRepetitionmsg3R17 struct {
	Value utils.ENUMERATED
}

const (
	BandnrPuschRepetitionmsg3R17EnumeratedNothing = iota
	BandnrPuschRepetitionmsg3R17EnumeratedSupported
)
