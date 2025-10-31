package ies

// RRCConnectionReject-criticalExtensions ::= CHOICE
const (
	RrcconnectionrejectCriticalextensionsChoiceNothing = iota
	RrcconnectionrejectCriticalextensionsChoiceC1
	RrcconnectionrejectCriticalextensionsChoiceCriticalextensionsfuture
)

type RrcconnectionrejectCriticalextensions struct {
	Choice                   uint64
	C1                       *RrcconnectionrejectCriticalextensionsC1
	Criticalextensionsfuture *RrcconnectionrejectCriticalextensionsCriticalextensionsfuture
}
