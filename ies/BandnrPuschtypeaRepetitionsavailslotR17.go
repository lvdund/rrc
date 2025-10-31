package ies

import "rrc/utils"

// BandNR-puschTypeA-RepetitionsAvailSlot-r17 ::= ENUMERATED
type BandnrPuschtypeaRepetitionsavailslotR17 struct {
	Value utils.ENUMERATED
}

const (
	BandnrPuschtypeaRepetitionsavailslotR17EnumeratedNothing = iota
	BandnrPuschtypeaRepetitionsavailslotR17EnumeratedSupported
)
