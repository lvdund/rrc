package ies

// RRCConnectionResume-NB-criticalExtensions ::= CHOICE
const (
	RrcconnectionresumeNbCriticalextensionsChoiceNothing = iota
	RrcconnectionresumeNbCriticalextensionsChoiceC1
	RrcconnectionresumeNbCriticalextensionsChoiceCriticalextensionsfuture
)

type RrcconnectionresumeNbCriticalextensions struct {
	Choice                   uint64
	C1                       *RrcconnectionresumeNbCriticalextensionsC1
	Criticalextensionsfuture *RrcconnectionresumeNbCriticalextensionsCriticalextensionsfuture
}
