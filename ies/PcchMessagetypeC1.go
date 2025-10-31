package ies

// PCCH-MessageType-c1 ::= CHOICE
const (
	PcchMessagetypeC1ChoiceNothing = iota
	PcchMessagetypeC1ChoicePaging
)

type PcchMessagetypeC1 struct {
	Choice uint64
	Paging *Paging
}
