package ies

import "rrc/utils"

// PUCCH-ResourceExt-v1610-pucch-RepetitionNrofSlots-r17 ::= ENUMERATED
type PucchResourceextV1610PucchRepetitionnrofslotsR17 struct {
	Value utils.ENUMERATED
}

const (
	PucchResourceextV1610PucchRepetitionnrofslotsR17EnumeratedNothing = iota
	PucchResourceextV1610PucchRepetitionnrofslotsR17EnumeratedN1
	PucchResourceextV1610PucchRepetitionnrofslotsR17EnumeratedN2
	PucchResourceextV1610PucchRepetitionnrofslotsR17EnumeratedN4
	PucchResourceextV1610PucchRepetitionnrofslotsR17EnumeratedN8
)
