package ies

// RRCConnectionReconfiguration-criticalExtensions-c1 ::= CHOICE
const (
	RrcconnectionreconfigurationCriticalextensionsC1ChoiceNothing = iota
	RrcconnectionreconfigurationCriticalextensionsC1ChoiceRrcconnectionreconfigurationR8
	RrcconnectionreconfigurationCriticalextensionsC1ChoiceSpare7
	RrcconnectionreconfigurationCriticalextensionsC1ChoiceSpare6
	RrcconnectionreconfigurationCriticalextensionsC1ChoiceSpare5
	RrcconnectionreconfigurationCriticalextensionsC1ChoiceSpare4
	RrcconnectionreconfigurationCriticalextensionsC1ChoiceSpare3
	RrcconnectionreconfigurationCriticalextensionsC1ChoiceSpare2
	RrcconnectionreconfigurationCriticalextensionsC1ChoiceSpare1
)

type RrcconnectionreconfigurationCriticalextensionsC1 struct {
	Choice                         uint64
	RrcconnectionreconfigurationR8 *RrcconnectionreconfigurationR8
	Spare7                         *struct{}
	Spare6                         *struct{}
	Spare5                         *struct{}
	Spare4                         *struct{}
	Spare3                         *struct{}
	Spare2                         *struct{}
	Spare1                         *struct{}
}
