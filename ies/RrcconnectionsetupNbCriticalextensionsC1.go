package ies

// RRCConnectionSetup-NB-criticalExtensions-c1 ::= CHOICE
const (
	RrcconnectionsetupNbCriticalextensionsC1ChoiceNothing = iota
	RrcconnectionsetupNbCriticalextensionsC1ChoiceRrcconnectionsetupR13
	RrcconnectionsetupNbCriticalextensionsC1ChoiceSpare1
)

type RrcconnectionsetupNbCriticalextensionsC1 struct {
	Choice                uint64
	RrcconnectionsetupR13 *RrcconnectionsetupNbR13
	Spare1                *struct{}
}
