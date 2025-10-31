package ies

// PUCCH-Resource-format ::= CHOICE
const (
	PucchResourceFormatChoiceNothing = iota
	PucchResourceFormatChoiceFormat0
	PucchResourceFormatChoiceFormat1
	PucchResourceFormatChoiceFormat2
	PucchResourceFormatChoiceFormat3
	PucchResourceFormatChoiceFormat4
)

type PucchResourceFormat struct {
	Choice  uint64
	Format0 *PucchFormat0
	Format1 *PucchFormat1
	Format2 *PucchFormat2
	Format3 *PucchFormat3
	Format4 *PucchFormat4
}
