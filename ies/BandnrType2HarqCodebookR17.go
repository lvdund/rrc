package ies

import "rrc/utils"

// BandNR-type2-HARQ-Codebook-r17 ::= ENUMERATED
type BandnrType2HarqCodebookR17 struct {
	Value utils.ENUMERATED
}

const (
	BandnrType2HarqCodebookR17EnumeratedNothing = iota
	BandnrType2HarqCodebookR17EnumeratedSupported
)
