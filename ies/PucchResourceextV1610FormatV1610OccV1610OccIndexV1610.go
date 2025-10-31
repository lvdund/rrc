package ies

import "rrc/utils"

// PUCCH-ResourceExt-v1610-format-v1610-occ-v1610-occ-Index-v1610 ::= ENUMERATED
type PucchResourceextV1610FormatV1610OccV1610OccIndexV1610 struct {
	Value utils.ENUMERATED
}

const (
	PucchResourceextV1610FormatV1610OccV1610OccIndexV1610EnumeratedNothing = iota
	PucchResourceextV1610FormatV1610OccV1610OccIndexV1610EnumeratedN0
	PucchResourceextV1610FormatV1610OccV1610OccIndexV1610EnumeratedN1
	PucchResourceextV1610FormatV1610OccV1610OccIndexV1610EnumeratedN2
	PucchResourceextV1610FormatV1610OccV1610OccIndexV1610EnumeratedN3
)
