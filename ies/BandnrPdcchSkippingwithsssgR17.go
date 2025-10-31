package ies

import "rrc/utils"

// BandNR-pdcch-SkippingWithSSSG-r17 ::= ENUMERATED
type BandnrPdcchSkippingwithsssgR17 struct {
	Value utils.ENUMERATED
}

const (
	BandnrPdcchSkippingwithsssgR17EnumeratedNothing = iota
	BandnrPdcchSkippingwithsssgR17EnumeratedSupported
)
