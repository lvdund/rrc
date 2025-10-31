package ies

import "rrc/utils"

// BandNR-bwp-DiffNumerology ::= ENUMERATED
type BandnrBwpDiffnumerology struct {
	Value utils.ENUMERATED
}

const (
	BandnrBwpDiffnumerologyEnumeratedNothing = iota
	BandnrBwpDiffnumerologyEnumeratedUpto4
)
