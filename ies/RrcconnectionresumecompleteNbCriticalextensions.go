package ies

// RRCConnectionResumeComplete-NB-criticalExtensions ::= CHOICE
const (
	RrcconnectionresumecompleteNbCriticalextensionsChoiceNothing = iota
	RrcconnectionresumecompleteNbCriticalextensionsChoiceRrcconnectionresumecompleteR13
	RrcconnectionresumecompleteNbCriticalextensionsChoiceCriticalextensionsfuture
)

type RrcconnectionresumecompleteNbCriticalextensions struct {
	Choice                         uint64
	RrcconnectionresumecompleteR13 *RrcconnectionresumecompleteNbR13
	Criticalextensionsfuture       *RrcconnectionresumecompleteNbCriticalextensionsCriticalextensionsfuture
}
