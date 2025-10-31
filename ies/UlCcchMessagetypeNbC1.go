package ies

// UL-CCCH-MessageType-NB-c1 ::= CHOICE
const (
	UlCcchMessagetypeNbC1ChoiceNothing = iota
	UlCcchMessagetypeNbC1ChoiceRrcconnectionreestablishmentrequestR13
	UlCcchMessagetypeNbC1ChoiceRrcconnectionrequestR13
	UlCcchMessagetypeNbC1ChoiceRrcconnectionresumerequestR13
	UlCcchMessagetypeNbC1ChoiceRrcearlydatarequestR15
)

type UlCcchMessagetypeNbC1 struct {
	Choice                                 uint64
	RrcconnectionreestablishmentrequestR13 *RrcconnectionreestablishmentrequestNb
	RrcconnectionrequestR13                *RrcconnectionrequestNb
	RrcconnectionresumerequestR13          *RrcconnectionresumerequestNb
	RrcearlydatarequestR15                 *RrcearlydatarequestNbR15
}
