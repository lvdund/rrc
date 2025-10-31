package ies

// RRCConnectionResume-NB-criticalExtensions-c1 ::= CHOICE
const (
	RrcconnectionresumeNbCriticalextensionsC1ChoiceNothing = iota
	RrcconnectionresumeNbCriticalextensionsC1ChoiceRrcconnectionresumeR13
	RrcconnectionresumeNbCriticalextensionsC1ChoiceSpare1
)

type RrcconnectionresumeNbCriticalextensionsC1 struct {
	Choice                 uint64
	RrcconnectionresumeR13 *RrcconnectionresumeNbR13
	Spare1                 *struct{}
}
