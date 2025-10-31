package ies

// RRCConnectionReject-NB-criticalExtensions-c1 ::= CHOICE
const (
	RrcconnectionrejectNbCriticalextensionsC1ChoiceNothing = iota
	RrcconnectionrejectNbCriticalextensionsC1ChoiceRrcconnectionrejectR13
	RrcconnectionrejectNbCriticalextensionsC1ChoiceSpare1
)

type RrcconnectionrejectNbCriticalextensionsC1 struct {
	Choice                 uint64
	RrcconnectionrejectR13 *RrcconnectionrejectNbR13
	Spare1                 *struct{}
}
