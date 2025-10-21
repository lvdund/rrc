package ies

import "rrc/utils"

// DeltaFList-PUCCH ::= SEQUENCE
type DeltaflistPucch struct {
	DeltafPucchFormat1  utils.ENUMERATED
	DeltafPucchFormat1b utils.ENUMERATED
	DeltafPucchFormat2  utils.ENUMERATED
	DeltafPucchFormat2a utils.ENUMERATED
	DeltafPucchFormat2b utils.ENUMERATED
}
