package ies

// DL-CCCH-MessageType-c1 ::= CHOICE
const (
	DlCcchMessagetypeC1ChoiceNothing = iota
	DlCcchMessagetypeC1ChoiceRrcconnectionreestablishment
	DlCcchMessagetypeC1ChoiceRrcconnectionreestablishmentreject
	DlCcchMessagetypeC1ChoiceRrcconnectionreject
	DlCcchMessagetypeC1ChoiceRrcconnectionsetup
)

type DlCcchMessagetypeC1 struct {
	Choice                             uint64
	Rrcconnectionreestablishment       *Rrcconnectionreestablishment
	Rrcconnectionreestablishmentreject *Rrcconnectionreestablishmentreject
	Rrcconnectionreject                *Rrcconnectionreject
	Rrcconnectionsetup                 *Rrcconnectionsetup
}
