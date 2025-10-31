package ies

import "rrc/utils"

// DeltaFList-PUCCH-deltaF-PUCCH-Format2a ::= ENUMERATED
type DeltaflistPucchDeltafPucchFormat2a struct {
	Value utils.ENUMERATED
}

const (
	DeltaflistPucchDeltafPucchFormat2aEnumeratedNothing = iota
	DeltaflistPucchDeltafPucchFormat2aEnumeratedDeltaf_2
	DeltaflistPucchDeltafPucchFormat2aEnumeratedDeltaf0
	DeltaflistPucchDeltafPucchFormat2aEnumeratedDeltaf2
)
