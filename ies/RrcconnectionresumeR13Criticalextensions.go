package ies

// RRCConnectionResume-r13-criticalExtensions ::= CHOICE
const (
	RrcconnectionresumeR13CriticalextensionsChoiceNothing = iota
	RrcconnectionresumeR13CriticalextensionsChoiceC1
	RrcconnectionresumeR13CriticalextensionsChoiceCriticalextensionsfuture
)

type RrcconnectionresumeR13Criticalextensions struct {
	Choice                   uint64
	C1                       *RrcconnectionresumeR13CriticalextensionsC1
	Criticalextensionsfuture *RrcconnectionresumeR13CriticalextensionsCriticalextensionsfuture
}
