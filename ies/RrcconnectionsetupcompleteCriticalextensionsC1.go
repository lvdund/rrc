package ies

// RRCConnectionSetupComplete-criticalExtensions-c1 ::= CHOICE
const (
	RrcconnectionsetupcompleteCriticalextensionsC1ChoiceNothing = iota
	RrcconnectionsetupcompleteCriticalextensionsC1ChoiceRrcconnectionsetupcompleteR8
	RrcconnectionsetupcompleteCriticalextensionsC1ChoiceSpare3
	RrcconnectionsetupcompleteCriticalextensionsC1ChoiceSpare2
	RrcconnectionsetupcompleteCriticalextensionsC1ChoiceSpare1
)

type RrcconnectionsetupcompleteCriticalextensionsC1 struct {
	Choice                       uint64
	RrcconnectionsetupcompleteR8 *RrcconnectionsetupcompleteR8
	Spare3                       *struct{}
	Spare2                       *struct{}
	Spare1                       *struct{}
}
