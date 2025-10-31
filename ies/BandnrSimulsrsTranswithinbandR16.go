package ies

import "rrc/utils"

// BandNR-simulSRS-TransWithinBand-r16 ::= ENUMERATED
type BandnrSimulsrsTranswithinbandR16 struct {
	Value utils.ENUMERATED
}

const (
	BandnrSimulsrsTranswithinbandR16EnumeratedNothing = iota
	BandnrSimulsrsTranswithinbandR16EnumeratedN2
)
