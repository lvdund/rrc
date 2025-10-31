package ies

import "rrc/utils"

// PUCCH-ResourceExt-v1610-format-v1610 ::= CHOICE
const (
	PucchResourceextV1610FormatV1610ChoiceNothing = iota
	PucchResourceextV1610FormatV1610ChoiceInterlace1V1610
	PucchResourceextV1610FormatV1610ChoiceOccV1610
)

type PucchResourceextV1610FormatV1610 struct {
	Choice          uint64
	Interlace1V1610 *utils.INTEGER `lb:0,ub:9`
	OccV1610        *PucchResourceextV1610FormatV1610OccV1610
}
