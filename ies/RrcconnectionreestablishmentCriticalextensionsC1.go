package ies

// RRCConnectionReestablishment-criticalExtensions-c1 ::= CHOICE
const (
	RrcconnectionreestablishmentCriticalextensionsC1ChoiceNothing = iota
	RrcconnectionreestablishmentCriticalextensionsC1ChoiceRrcconnectionreestablishmentR8
	RrcconnectionreestablishmentCriticalextensionsC1ChoiceSpare7
	RrcconnectionreestablishmentCriticalextensionsC1ChoiceSpare6
	RrcconnectionreestablishmentCriticalextensionsC1ChoiceSpare5
	RrcconnectionreestablishmentCriticalextensionsC1ChoiceSpare4
	RrcconnectionreestablishmentCriticalextensionsC1ChoiceSpare3
	RrcconnectionreestablishmentCriticalextensionsC1ChoiceSpare2
	RrcconnectionreestablishmentCriticalextensionsC1ChoiceSpare1
)

type RrcconnectionreestablishmentCriticalextensionsC1 struct {
	Choice                         uint64
	RrcconnectionreestablishmentR8 *RrcconnectionreestablishmentR8
	Spare7                         *struct{}
	Spare6                         *struct{}
	Spare5                         *struct{}
	Spare4                         *struct{}
	Spare3                         *struct{}
	Spare2                         *struct{}
	Spare1                         *struct{}
}
