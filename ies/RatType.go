package ies

import "rrc/utils"

// RAT-Type ::= utils.ENUMERATED // Extensible
type RatType struct {
	Value utils.ENUMERATED
}

const (
	RatTypeEnumeratedNothing = iota
	RatTypeEnumeratedNr
	RatTypeEnumeratedEutra_Nr
	RatTypeEnumeratedEutra
	RatTypeEnumeratedUtra_Fdd_V1610
)
