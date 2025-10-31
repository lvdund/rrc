package ies

// MCCH-MessageType-c1 ::= CHOICE
const (
	McchMessagetypeC1ChoiceNothing = iota
	McchMessagetypeC1ChoiceMbsfnareaconfigurationR9
)

type McchMessagetypeC1 struct {
	Choice                   uint64
	MbsfnareaconfigurationR9 *MbsfnareaconfigurationR9
}
