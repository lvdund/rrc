package ies

import "rrc/utils"

// DeltaFList-PUCCH-deltaF-PUCCH-Format2b ::= ENUMERATED
type DeltaflistPucchDeltafPucchFormat2b struct {
	Value utils.ENUMERATED
}

const (
	DeltaflistPucchDeltafPucchFormat2bEnumeratedNothing = iota
	DeltaflistPucchDeltafPucchFormat2bEnumeratedDeltaf_2
	DeltaflistPucchDeltafPucchFormat2bEnumeratedDeltaf0
	DeltaflistPucchDeltafPucchFormat2bEnumeratedDeltaf2
)
