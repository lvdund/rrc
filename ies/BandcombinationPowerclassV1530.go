package ies

import "rrc/utils"

// BandCombination-powerClass-v1530 ::= ENUMERATED
type BandcombinationPowerclassV1530 struct {
	Value utils.ENUMERATED
}

const (
	BandcombinationPowerclassV1530EnumeratedNothing = iota
	BandcombinationPowerclassV1530EnumeratedPc2
)
