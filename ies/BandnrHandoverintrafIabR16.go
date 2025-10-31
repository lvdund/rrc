package ies

import "rrc/utils"

// BandNR-handoverIntraF-IAB-r16 ::= ENUMERATED
type BandnrHandoverintrafIabR16 struct {
	Value utils.ENUMERATED
}

const (
	BandnrHandoverintrafIabR16EnumeratedNothing = iota
	BandnrHandoverintrafIabR16EnumeratedSupported
)
