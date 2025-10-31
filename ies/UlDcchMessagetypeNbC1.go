package ies

// UL-DCCH-MessageType-NB-c1 ::= CHOICE
const (
	UlDcchMessagetypeNbC1ChoiceNothing = iota
	UlDcchMessagetypeNbC1ChoiceRrcconnectionreconfigurationcompleteR13
	UlDcchMessagetypeNbC1ChoiceRrcconnectionreestablishmentcompleteR13
	UlDcchMessagetypeNbC1ChoiceRrcconnectionsetupcompleteR13
	UlDcchMessagetypeNbC1ChoiceSecuritymodecompleteR13
	UlDcchMessagetypeNbC1ChoiceSecuritymodefailureR13
	UlDcchMessagetypeNbC1ChoiceUecapabilityinformationR13
	UlDcchMessagetypeNbC1ChoiceUlinformationtransferR13
	UlDcchMessagetypeNbC1ChoiceRrcconnectionresumecompleteR13
	UlDcchMessagetypeNbC1ChoiceUeinformationresponseR16
	UlDcchMessagetypeNbC1ChoicePurconfigurationrequestR16
	UlDcchMessagetypeNbC1ChoiceSpare6
	UlDcchMessagetypeNbC1ChoiceSpare5
	UlDcchMessagetypeNbC1ChoiceSpare4
	UlDcchMessagetypeNbC1ChoiceSpare3
	UlDcchMessagetypeNbC1ChoiceSpare2
	UlDcchMessagetypeNbC1ChoiceSpare1
)

type UlDcchMessagetypeNbC1 struct {
	Choice                                  uint64
	RrcconnectionreconfigurationcompleteR13 *RrcconnectionreconfigurationcompleteNb
	RrcconnectionreestablishmentcompleteR13 *RrcconnectionreestablishmentcompleteNb
	RrcconnectionsetupcompleteR13           *RrcconnectionsetupcompleteNb
	SecuritymodecompleteR13                 *Securitymodecomplete
	SecuritymodefailureR13                  *Securitymodefailure
	UecapabilityinformationR13              *UecapabilityinformationNb
	UlinformationtransferR13                *UlinformationtransferNb
	RrcconnectionresumecompleteR13          *RrcconnectionresumecompleteNb
	UeinformationresponseR16                *UeinformationresponseNbR16
	PurconfigurationrequestR16              *PurconfigurationrequestNbR16
	Spare6                                  *struct{}
	Spare5                                  *struct{}
	Spare4                                  *struct{}
	Spare3                                  *struct{}
	Spare2                                  *struct{}
	Spare1                                  *struct{}
}
