package ies

import "rrc/utils"

// BandCombination-v1610-powerClassNRPart-r16 ::= ENUMERATED
type BandcombinationV1610PowerclassnrpartR16 struct {
	Value utils.ENUMERATED
}

const (
	BandcombinationV1610PowerclassnrpartR16EnumeratedNothing = iota
	BandcombinationV1610PowerclassnrpartR16EnumeratedPc1
	BandcombinationV1610PowerclassnrpartR16EnumeratedPc2
	BandcombinationV1610PowerclassnrpartR16EnumeratedPc3
	BandcombinationV1610PowerclassnrpartR16EnumeratedPc5
)
