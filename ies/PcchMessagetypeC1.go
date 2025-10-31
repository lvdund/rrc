package ies

// PCCH-MessageType-c1 ::= CHOICE
const (
	PcchMessagetypeC1ChoiceNothing = iota
	PcchMessagetypeC1ChoicePaging
	PcchMessagetypeC1ChoiceSpare1
)

type PcchMessagetypeC1 struct {
	Choice uint64
	Paging *Paging
	Spare1 *struct{}
}
