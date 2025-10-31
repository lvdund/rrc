package ies

import "rrc/utils"

// BandNR-rateMatchingLTE-CRS ::= ENUMERATED
type BandnrRatematchinglteCrs struct {
	Value utils.ENUMERATED
}

const (
	BandnrRatematchinglteCrsEnumeratedNothing = iota
	BandnrRatematchinglteCrsEnumeratedSupported
)
