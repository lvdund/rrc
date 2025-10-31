package ies

import "rrc/utils"

// BandCombination-v1610-powerClass-v1610 ::= ENUMERATED
type BandcombinationV1610PowerclassV1610 struct {
	Value utils.ENUMERATED
}

const (
	BandcombinationV1610PowerclassV1610EnumeratedNothing = iota
	BandcombinationV1610PowerclassV1610EnumeratedPc1dot5
)
