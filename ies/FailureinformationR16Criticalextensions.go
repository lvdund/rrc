package ies

// FailureInformation-r16-criticalExtensions ::= CHOICE
const (
	FailureinformationR16CriticalextensionsChoiceNothing = iota
	FailureinformationR16CriticalextensionsChoiceFailureinformationR16
	FailureinformationR16CriticalextensionsChoiceCriticalextensionsfuture
)

type FailureinformationR16Criticalextensions struct {
	Choice                   uint64
	FailureinformationR16    *FailureinformationR16
	Criticalextensionsfuture *FailureinformationR16CriticalextensionsCriticalextensionsfuture
}
