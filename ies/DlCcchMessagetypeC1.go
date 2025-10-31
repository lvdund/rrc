package ies

// DL-CCCH-MessageType-c1 ::= CHOICE
const (
	DlCcchMessagetypeC1ChoiceNothing = iota
	DlCcchMessagetypeC1ChoiceRrcreject
	DlCcchMessagetypeC1ChoiceRrcsetup
	DlCcchMessagetypeC1ChoiceSpare2
	DlCcchMessagetypeC1ChoiceSpare1
)

type DlCcchMessagetypeC1 struct {
	Choice    uint64
	Rrcreject *Rrcreject
	Rrcsetup  *Rrcsetup
	Spare2    *struct{}
	Spare1    *struct{}
}
