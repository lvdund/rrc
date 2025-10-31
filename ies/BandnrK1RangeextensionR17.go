package ies

import "rrc/utils"

// BandNR-k1-RangeExtension-r17 ::= ENUMERATED
type BandnrK1RangeextensionR17 struct {
	Value utils.ENUMERATED
}

const (
	BandnrK1RangeextensionR17EnumeratedNothing = iota
	BandnrK1RangeextensionR17EnumeratedSupported
)
