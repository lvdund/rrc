package ies

// HandoverPreparationInformation-NB-criticalExtensions ::= CHOICE
const (
	HandoverpreparationinformationNbCriticalextensionsChoiceNothing = iota
	HandoverpreparationinformationNbCriticalextensionsChoiceC1
	HandoverpreparationinformationNbCriticalextensionsChoiceCriticalextensionsfuture
)

type HandoverpreparationinformationNbCriticalextensions struct {
	Choice                   uint64
	C1                       *HandoverpreparationinformationNbCriticalextensionsC1
	Criticalextensionsfuture *HandoverpreparationinformationNbCriticalextensionsCriticalextensionsfuture
}
