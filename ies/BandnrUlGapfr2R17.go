package ies

import "rrc/utils"

// BandNR-ul-GapFR2-r17 ::= ENUMERATED
type BandnrUlGapfr2R17 struct {
	Value utils.ENUMERATED
}

const (
	BandnrUlGapfr2R17EnumeratedNothing = iota
	BandnrUlGapfr2R17EnumeratedSupported
)
