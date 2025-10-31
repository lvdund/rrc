package ies

// RRCConnectionSetup-NB-criticalExtensions ::= CHOICE
const (
	RrcconnectionsetupNbCriticalextensionsChoiceNothing = iota
	RrcconnectionsetupNbCriticalextensionsChoiceC1
	RrcconnectionsetupNbCriticalextensionsChoiceCriticalextensionsfuture
)

type RrcconnectionsetupNbCriticalextensions struct {
	Choice                   uint64
	C1                       *RrcconnectionsetupNbCriticalextensionsC1
	Criticalextensionsfuture *RrcconnectionsetupNbCriticalextensionsCriticalextensionsfuture
}
