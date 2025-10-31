package ies

// RRCConnectionReject-NB-criticalExtensions ::= CHOICE
const (
	RrcconnectionrejectNbCriticalextensionsChoiceNothing = iota
	RrcconnectionrejectNbCriticalextensionsChoiceC1
	RrcconnectionrejectNbCriticalextensionsChoiceCriticalextensionsfuture
)

type RrcconnectionrejectNbCriticalextensions struct {
	Choice                   uint64
	C1                       *RrcconnectionrejectNbCriticalextensionsC1
	Criticalextensionsfuture *RrcconnectionrejectNbCriticalextensionsCriticalextensionsfuture
}
