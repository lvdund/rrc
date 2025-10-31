package ies

// RRCConnectionSetupComplete-criticalExtensions ::= CHOICE
const (
	RrcconnectionsetupcompleteCriticalextensionsChoiceNothing = iota
	RrcconnectionsetupcompleteCriticalextensionsChoiceC1
	RrcconnectionsetupcompleteCriticalextensionsChoiceCriticalextensionsfuture
)

type RrcconnectionsetupcompleteCriticalextensions struct {
	Choice                   uint64
	C1                       *RrcconnectionsetupcompleteCriticalextensionsC1
	Criticalextensionsfuture *RrcconnectionsetupcompleteCriticalextensionsCriticalextensionsfuture
}
