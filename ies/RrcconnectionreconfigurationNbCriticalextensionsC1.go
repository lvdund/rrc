package ies

// RRCConnectionReconfiguration-NB-criticalExtensions-c1 ::= CHOICE
const (
	RrcconnectionreconfigurationNbCriticalextensionsC1ChoiceNothing = iota
	RrcconnectionreconfigurationNbCriticalextensionsC1ChoiceRrcconnectionreconfigurationR13
	RrcconnectionreconfigurationNbCriticalextensionsC1ChoiceSpare1
)

type RrcconnectionreconfigurationNbCriticalextensionsC1 struct {
	Choice                          uint64
	RrcconnectionreconfigurationR13 *RrcconnectionreconfigurationNbR13
	Spare1                          *struct{}
}
