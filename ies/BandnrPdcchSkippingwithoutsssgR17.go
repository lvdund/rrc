package ies

import "rrc/utils"

// BandNR-pdcch-SkippingWithoutSSSG-r17 ::= ENUMERATED
type BandnrPdcchSkippingwithoutsssgR17 struct {
	Value utils.ENUMERATED
}

const (
	BandnrPdcchSkippingwithoutsssgR17EnumeratedNothing = iota
	BandnrPdcchSkippingwithoutsssgR17EnumeratedSupported
)
