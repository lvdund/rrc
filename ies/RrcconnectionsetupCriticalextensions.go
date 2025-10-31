package ies

// RRCConnectionSetup-criticalExtensions ::= CHOICE
const (
	RrcconnectionsetupCriticalextensionsChoiceNothing = iota
	RrcconnectionsetupCriticalextensionsChoiceC1
	RrcconnectionsetupCriticalextensionsChoiceCriticalextensionsfuture
)

type RrcconnectionsetupCriticalextensions struct {
	Choice                   uint64
	C1                       *RrcconnectionsetupCriticalextensionsC1
	Criticalextensionsfuture *RrcconnectionsetupCriticalextensionsCriticalextensionsfuture
}
