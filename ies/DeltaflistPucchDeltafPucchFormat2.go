package ies

import "rrc/utils"

// DeltaFList-PUCCH-deltaF-PUCCH-Format2 ::= ENUMERATED
type DeltaflistPucchDeltafPucchFormat2 struct {
	Value utils.ENUMERATED
}

const (
	DeltaflistPucchDeltafPucchFormat2EnumeratedNothing = iota
	DeltaflistPucchDeltafPucchFormat2EnumeratedDeltaf_2
	DeltaflistPucchDeltafPucchFormat2EnumeratedDeltaf0
	DeltaflistPucchDeltafPucchFormat2EnumeratedDeltaf1
	DeltaflistPucchDeltafPucchFormat2EnumeratedDeltaf2
)
