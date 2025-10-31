package ies

// HandoverFromEUTRAPreparationRequest-criticalExtensions ::= CHOICE
const (
	HandoverfromeutrapreparationrequestCriticalextensionsChoiceNothing = iota
	HandoverfromeutrapreparationrequestCriticalextensionsChoiceC1
	HandoverfromeutrapreparationrequestCriticalextensionsChoiceCriticalextensionsfuture
)

type HandoverfromeutrapreparationrequestCriticalextensions struct {
	Choice                   uint64
	C1                       *HandoverfromeutrapreparationrequestCriticalextensionsC1
	Criticalextensionsfuture *HandoverfromeutrapreparationrequestCriticalextensionsCriticalextensionsfuture
}
