package ies

import "rrc/utils"

// DeltaFList-PUCCH-deltaF-PUCCH-Format1b ::= ENUMERATED
type DeltaflistPucchDeltafPucchFormat1b struct {
	Value utils.ENUMERATED
}

const (
	DeltaflistPucchDeltafPucchFormat1bEnumeratedNothing = iota
	DeltaflistPucchDeltafPucchFormat1bEnumeratedDeltaf1
	DeltaflistPucchDeltafPucchFormat1bEnumeratedDeltaf3
	DeltaflistPucchDeltafPucchFormat1bEnumeratedDeltaf5
)
