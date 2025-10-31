package ies

import "rrc/utils"

// BandNR-rasterShift7dot5-IAB-r16 ::= ENUMERATED
type BandnrRastershift7dot5IabR16 struct {
	Value utils.ENUMERATED
}

const (
	BandnrRastershift7dot5IabR16EnumeratedNothing = iota
	BandnrRastershift7dot5IabR16EnumeratedSupported
)
