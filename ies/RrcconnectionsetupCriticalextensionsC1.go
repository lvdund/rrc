package ies

// RRCConnectionSetup-criticalExtensions-c1 ::= CHOICE
const (
	RrcconnectionsetupCriticalextensionsC1ChoiceNothing = iota
	RrcconnectionsetupCriticalextensionsC1ChoiceRrcconnectionsetupR8
	RrcconnectionsetupCriticalextensionsC1ChoiceSpare7
	RrcconnectionsetupCriticalextensionsC1ChoiceSpare6
	RrcconnectionsetupCriticalextensionsC1ChoiceSpare5
	RrcconnectionsetupCriticalextensionsC1ChoiceSpare4
	RrcconnectionsetupCriticalextensionsC1ChoiceSpare3
	RrcconnectionsetupCriticalextensionsC1ChoiceSpare2
	RrcconnectionsetupCriticalextensionsC1ChoiceSpare1
)

type RrcconnectionsetupCriticalextensionsC1 struct {
	Choice               uint64
	RrcconnectionsetupR8 *RrcconnectionsetupR8
	Spare7               *struct{}
	Spare6               *struct{}
	Spare5               *struct{}
	Spare4               *struct{}
	Spare3               *struct{}
	Spare2               *struct{}
	Spare1               *struct{}
}
