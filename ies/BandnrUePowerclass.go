package ies

import "rrc/utils"

// BandNR-ue-PowerClass ::= ENUMERATED
type BandnrUePowerclass struct {
	Value utils.ENUMERATED
}

const (
	BandnrUePowerclassEnumeratedNothing = iota
	BandnrUePowerclassEnumeratedPc1
	BandnrUePowerclassEnumeratedPc2
	BandnrUePowerclassEnumeratedPc3
	BandnrUePowerclassEnumeratedPc4
)
