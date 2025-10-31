package ies

import "rrc/utils"

// RAT-Type ::= utils.ENUMERATED // Extensible
type RatType struct {
	Value utils.ENUMERATED
}

const (
	RatTypeEnumeratedNothing = iota
	RatTypeEnumeratedEutra
	RatTypeEnumeratedUtra
	RatTypeEnumeratedGeran_Cs
	RatTypeEnumeratedGeran_Ps
	RatTypeEnumeratedCdma2000_1xrtt
	RatTypeEnumeratedNr
	RatTypeEnumeratedEutra_Nr
	RatTypeEnumeratedSpare1
)
