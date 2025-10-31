package ies

import "rrc/utils"

// BandNR-type1-HARQ-Codebook-r17 ::= ENUMERATED
type BandnrType1HarqCodebookR17 struct {
	Value utils.ENUMERATED
}

const (
	BandnrType1HarqCodebookR17EnumeratedNothing = iota
	BandnrType1HarqCodebookR17EnumeratedSupported
)
