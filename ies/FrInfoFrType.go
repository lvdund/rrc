package ies

import "rrc/utils"

// FR-Info-fr-Type ::= ENUMERATED
type FrInfoFrType struct {
	Value utils.ENUMERATED
}

const (
	FrInfoFrTypeEnumeratedNothing = iota
	FrInfoFrTypeEnumeratedFr1
	FrInfoFrTypeEnumeratedFr2
)
