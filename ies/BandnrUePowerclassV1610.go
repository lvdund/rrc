package ies

import "rrc/utils"

// BandNR-ue-PowerClass-v1610 ::= ENUMERATED
type BandnrUePowerclassV1610 struct {
	Value utils.ENUMERATED
}

const (
	BandnrUePowerclassV1610EnumeratedNothing = iota
	BandnrUePowerclassV1610EnumeratedPc1dot5
)
