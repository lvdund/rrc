package ies

import "rrc/utils"

// BandNR-overlapRateMatchingEUTRA-CRS-r16 ::= ENUMERATED
type BandnrOverlapratematchingeutraCrsR16 struct {
	Value utils.ENUMERATED
}

const (
	BandnrOverlapratematchingeutraCrsR16EnumeratedNothing = iota
	BandnrOverlapratematchingeutraCrsR16EnumeratedSupported
)
