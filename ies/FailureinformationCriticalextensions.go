package ies

// FailureInformation-criticalExtensions ::= CHOICE
const (
	FailureinformationCriticalextensionsChoiceNothing = iota
	FailureinformationCriticalextensionsChoiceFailureinformation
	FailureinformationCriticalextensionsChoiceCriticalextensionsfuture
)

type FailureinformationCriticalextensions struct {
	Choice                   uint64
	Failureinformation       *Failureinformation
	Criticalextensionsfuture *FailureinformationCriticalextensionsCriticalextensionsfuture
}
