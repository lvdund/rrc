package ies

import "rrc/utils"

// BandNR-bwp-SameNumerology ::= ENUMERATED
type BandnrBwpSamenumerology struct {
	Value utils.ENUMERATED
}

const (
	BandnrBwpSamenumerologyEnumeratedNothing = iota
	BandnrBwpSamenumerologyEnumeratedUpto2
	BandnrBwpSamenumerologyEnumeratedUpto4
)
