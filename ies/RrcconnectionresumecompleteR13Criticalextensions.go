package ies

// RRCConnectionResumeComplete-r13-criticalExtensions ::= CHOICE
const (
	RrcconnectionresumecompleteR13CriticalextensionsChoiceNothing = iota
	RrcconnectionresumecompleteR13CriticalextensionsChoiceRrcconnectionresumecompleteR13
	RrcconnectionresumecompleteR13CriticalextensionsChoiceCriticalextensionsfuture
)

type RrcconnectionresumecompleteR13Criticalextensions struct {
	Choice                         uint64
	RrcconnectionresumecompleteR13 *RrcconnectionresumecompleteR13
	Criticalextensionsfuture       *RrcconnectionresumecompleteR13CriticalextensionsCriticalextensionsfuture
}
