package ies

// PUCCH-ResourceExt-v1610 ::= SEQUENCE
// Extensible
type PucchResourceextV1610 struct {
	InterlaceallocationR16      *PucchResourceextV1610InterlaceallocationR16
	FormatV1610                 *PucchResourceextV1610FormatV1610
	FormatextV1700              *PucchResourceextV1610FormatextV1700              `ext`
	PucchRepetitionnrofslotsR17 *PucchResourceextV1610PucchRepetitionnrofslotsR17 `ext`
}
