package ies

// HandoverPreparationInformation-criticalExtensions ::= CHOICE
const (
	HandoverpreparationinformationCriticalextensionsChoiceNothing = iota
	HandoverpreparationinformationCriticalextensionsChoiceC1
	HandoverpreparationinformationCriticalextensionsChoiceCriticalextensionsfuture
)

type HandoverpreparationinformationCriticalextensions struct {
	Choice                   uint64
	C1                       *HandoverpreparationinformationCriticalextensionsC1
	Criticalextensionsfuture *HandoverpreparationinformationCriticalextensionsCriticalextensionsfuture
}
