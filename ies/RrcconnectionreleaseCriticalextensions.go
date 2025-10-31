package ies

// RRCConnectionRelease-criticalExtensions ::= CHOICE
const (
	RrcconnectionreleaseCriticalextensionsChoiceNothing = iota
	RrcconnectionreleaseCriticalextensionsChoiceC1
	RrcconnectionreleaseCriticalextensionsChoiceCriticalextensionsfuture
)

type RrcconnectionreleaseCriticalextensions struct {
	Choice                   uint64
	C1                       *RrcconnectionreleaseCriticalextensionsC1
	Criticalextensionsfuture *RrcconnectionreleaseCriticalextensionsCriticalextensionsfuture
}
