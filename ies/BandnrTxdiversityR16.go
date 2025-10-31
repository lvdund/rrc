package ies

import "rrc/utils"

// BandNR-txDiversity-r16 ::= ENUMERATED
type BandnrTxdiversityR16 struct {
	Value utils.ENUMERATED
}

const (
	BandnrTxdiversityR16EnumeratedNothing = iota
	BandnrTxdiversityR16EnumeratedSupported
)
