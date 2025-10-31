package ies

// RRCConnectionReject-criticalExtensions-c1 ::= CHOICE
const (
	RrcconnectionrejectCriticalextensionsC1ChoiceNothing = iota
	RrcconnectionrejectCriticalextensionsC1ChoiceRrcconnectionrejectR8
	RrcconnectionrejectCriticalextensionsC1ChoiceSpare3
	RrcconnectionrejectCriticalextensionsC1ChoiceSpare2
	RrcconnectionrejectCriticalextensionsC1ChoiceSpare1
)

type RrcconnectionrejectCriticalextensionsC1 struct {
	Choice                uint64
	RrcconnectionrejectR8 *RrcconnectionrejectR8
	Spare3                *struct{}
	Spare2                *struct{}
	Spare1                *struct{}
}
