package ies

// RRCConnectionResume-r13-criticalExtensions-c1 ::= CHOICE
const (
	RrcconnectionresumeR13CriticalextensionsC1ChoiceNothing = iota
	RrcconnectionresumeR13CriticalextensionsC1ChoiceRrcconnectionresumeR13
	RrcconnectionresumeR13CriticalextensionsC1ChoiceSpare3
	RrcconnectionresumeR13CriticalextensionsC1ChoiceSpare2
	RrcconnectionresumeR13CriticalextensionsC1ChoiceSpare1
)

type RrcconnectionresumeR13CriticalextensionsC1 struct {
	Choice                 uint64
	RrcconnectionresumeR13 *RrcconnectionresumeR13
	Spare3                 *struct{}
	Spare2                 *struct{}
	Spare1                 *struct{}
}
