package ies

// RRCConnectionReconfiguration-NB-criticalExtensions ::= CHOICE
const (
	RrcconnectionreconfigurationNbCriticalextensionsChoiceNothing = iota
	RrcconnectionreconfigurationNbCriticalextensionsChoiceC1
	RrcconnectionreconfigurationNbCriticalextensionsChoiceCriticalextensionsfuture
)

type RrcconnectionreconfigurationNbCriticalextensions struct {
	Choice                   uint64
	C1                       *RrcconnectionreconfigurationNbCriticalextensionsC1
	Criticalextensionsfuture *RrcconnectionreconfigurationNbCriticalextensionsCriticalextensionsfuture
}
