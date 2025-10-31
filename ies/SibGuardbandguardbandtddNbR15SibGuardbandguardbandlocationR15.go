package ies

import "rrc/utils"

// SIB-GuardbandGuardbandTDD-NB-r15-sib-GuardbandGuardbandLocation-r15 ::= ENUMERATED
type SibGuardbandguardbandtddNbR15SibGuardbandguardbandlocationR15 struct {
	Value utils.ENUMERATED
}

const (
	SibGuardbandguardbandtddNbR15SibGuardbandguardbandlocationR15EnumeratedNothing = iota
	SibGuardbandguardbandtddNbR15SibGuardbandguardbandlocationR15EnumeratedSame
	SibGuardbandguardbandtddNbR15SibGuardbandguardbandlocationR15EnumeratedOpposite
)
