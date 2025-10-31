package ies

// RRCConnectionRelease-NB-criticalExtensions ::= CHOICE
const (
	RrcconnectionreleaseNbCriticalextensionsChoiceNothing = iota
	RrcconnectionreleaseNbCriticalextensionsChoiceC1
	RrcconnectionreleaseNbCriticalextensionsChoiceCriticalextensionsfuture
)

type RrcconnectionreleaseNbCriticalextensions struct {
	Choice                   uint64
	C1                       *RrcconnectionreleaseNbCriticalextensionsC1
	Criticalextensionsfuture *RrcconnectionreleaseNbCriticalextensionsCriticalextensionsfuture
}
