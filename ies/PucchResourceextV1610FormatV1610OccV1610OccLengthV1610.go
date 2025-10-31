package ies

import "rrc/utils"

// PUCCH-ResourceExt-v1610-format-v1610-occ-v1610-occ-Length-v1610 ::= ENUMERATED
type PucchResourceextV1610FormatV1610OccV1610OccLengthV1610 struct {
	Value utils.ENUMERATED
}

const (
	PucchResourceextV1610FormatV1610OccV1610OccLengthV1610EnumeratedNothing = iota
	PucchResourceextV1610FormatV1610OccV1610OccLengthV1610EnumeratedN2
	PucchResourceextV1610FormatV1610OccV1610OccLengthV1610EnumeratedN4
)
