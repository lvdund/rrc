package ies

import "rrc/utils"

// BandNR-ue-PowerClass-v1700 ::= ENUMERATED
type BandnrUePowerclassV1700 struct {
	Value utils.ENUMERATED
}

const (
	BandnrUePowerclassV1700EnumeratedNothing = iota
	BandnrUePowerclassV1700EnumeratedPc5
	BandnrUePowerclassV1700EnumeratedPc6
	BandnrUePowerclassV1700EnumeratedPc7
)
