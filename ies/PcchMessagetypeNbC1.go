package ies

// PCCH-MessageType-NB-c1 ::= CHOICE
const (
	PcchMessagetypeNbC1ChoiceNothing = iota
	PcchMessagetypeNbC1ChoicePagingR13
)

type PcchMessagetypeNbC1 struct {
	Choice    uint64
	PagingR13 *PagingNb
}
