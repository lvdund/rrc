package ies

// RRCConnectionReconfiguration-criticalExtensions ::= CHOICE
const (
	RrcconnectionreconfigurationCriticalextensionsChoiceNothing = iota
	RrcconnectionreconfigurationCriticalextensionsChoiceC1
	RrcconnectionreconfigurationCriticalextensionsChoiceCriticalextensionsfuture
)

type RrcconnectionreconfigurationCriticalextensions struct {
	Choice                   uint64
	C1                       *RrcconnectionreconfigurationCriticalextensionsC1
	Criticalextensionsfuture *RrcconnectionreconfigurationCriticalextensionsCriticalextensionsfuture
}
