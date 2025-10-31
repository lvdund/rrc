package ies

// DL-CCCH-MessageType-NB-c1 ::= CHOICE
const (
	DlCcchMessagetypeNbC1ChoiceNothing = iota
	DlCcchMessagetypeNbC1ChoiceRrcconnectionreestablishmentR13
	DlCcchMessagetypeNbC1ChoiceRrcconnectionreestablishmentrejectR13
	DlCcchMessagetypeNbC1ChoiceRrcconnectionrejectR13
	DlCcchMessagetypeNbC1ChoiceRrcconnectionsetupR13
	DlCcchMessagetypeNbC1ChoiceRrcearlydatacompleteR15
	DlCcchMessagetypeNbC1ChoiceSpare3
	DlCcchMessagetypeNbC1ChoiceSpare2
	DlCcchMessagetypeNbC1ChoiceSpare1
)

type DlCcchMessagetypeNbC1 struct {
	Choice                                uint64
	RrcconnectionreestablishmentR13       *RrcconnectionreestablishmentNb
	RrcconnectionreestablishmentrejectR13 *Rrcconnectionreestablishmentreject
	RrcconnectionrejectR13                *RrcconnectionrejectNb
	RrcconnectionsetupR13                 *RrcconnectionsetupNb
	RrcearlydatacompleteR15               *RrcearlydatacompleteNbR15
	Spare3                                *struct{}
	Spare2                                *struct{}
	Spare1                                *struct{}
}
