package ies

// RRCConnectionRelease-criticalExtensions-c1 ::= CHOICE
const (
	RrcconnectionreleaseCriticalextensionsC1ChoiceNothing = iota
	RrcconnectionreleaseCriticalextensionsC1ChoiceRrcconnectionreleaseR8
	RrcconnectionreleaseCriticalextensionsC1ChoiceSpare3
	RrcconnectionreleaseCriticalextensionsC1ChoiceSpare2
	RrcconnectionreleaseCriticalextensionsC1ChoiceSpare1
)

type RrcconnectionreleaseCriticalextensionsC1 struct {
	Choice                 uint64
	RrcconnectionreleaseR8 *RrcconnectionreleaseR8
	Spare3                 *struct{}
	Spare2                 *struct{}
	Spare1                 *struct{}
}
