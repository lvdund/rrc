package ies

// RRCConnectionReestablishment-NB-criticalExtensions-c1 ::= CHOICE
const (
	RrcconnectionreestablishmentNbCriticalextensionsC1ChoiceNothing = iota
	RrcconnectionreestablishmentNbCriticalextensionsC1ChoiceRrcconnectionreestablishmentR13
	RrcconnectionreestablishmentNbCriticalextensionsC1ChoiceSpare1
)

type RrcconnectionreestablishmentNbCriticalextensionsC1 struct {
	Choice                          uint64
	RrcconnectionreestablishmentR13 *RrcconnectionreestablishmentNbR13
	Spare1                          *struct{}
}
