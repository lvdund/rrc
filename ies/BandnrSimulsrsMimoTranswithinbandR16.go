package ies

import "rrc/utils"

// BandNR-simulSRS-MIMO-TransWithinBand-r16 ::= ENUMERATED
type BandnrSimulsrsMimoTranswithinbandR16 struct {
	Value utils.ENUMERATED
}

const (
	BandnrSimulsrsMimoTranswithinbandR16EnumeratedNothing = iota
	BandnrSimulsrsMimoTranswithinbandR16EnumeratedN2
)
