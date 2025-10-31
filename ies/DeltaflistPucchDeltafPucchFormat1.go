package ies

import "rrc/utils"

// DeltaFList-PUCCH-deltaF-PUCCH-Format1 ::= ENUMERATED
type DeltaflistPucchDeltafPucchFormat1 struct {
	Value utils.ENUMERATED
}

const (
	DeltaflistPucchDeltafPucchFormat1EnumeratedNothing = iota
	DeltaflistPucchDeltafPucchFormat1EnumeratedDeltaf_2
	DeltaflistPucchDeltafPucchFormat1EnumeratedDeltaf0
	DeltaflistPucchDeltafPucchFormat1EnumeratedDeltaf2
)
