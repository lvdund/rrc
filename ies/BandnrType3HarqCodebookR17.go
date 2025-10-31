package ies

import "rrc/utils"

// BandNR-type3-HARQ-Codebook-r17 ::= ENUMERATED
type BandnrType3HarqCodebookR17 struct {
	Value utils.ENUMERATED
}

const (
	BandnrType3HarqCodebookR17EnumeratedNothing = iota
	BandnrType3HarqCodebookR17EnumeratedSupported
)
