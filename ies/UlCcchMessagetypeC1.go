package ies

// UL-CCCH-MessageType-c1 ::= CHOICE
const (
	UlCcchMessagetypeC1ChoiceNothing = iota
	UlCcchMessagetypeC1ChoiceRrcconnectionreestablishmentrequest
	UlCcchMessagetypeC1ChoiceRrcconnectionrequest
)

type UlCcchMessagetypeC1 struct {
	Choice                              uint64
	Rrcconnectionreestablishmentrequest *Rrcconnectionreestablishmentrequest
	Rrcconnectionrequest                *Rrcconnectionrequest
}
