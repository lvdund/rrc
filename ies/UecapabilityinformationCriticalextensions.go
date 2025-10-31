package ies

// UECapabilityInformation-criticalExtensions ::= CHOICE
const (
	UecapabilityinformationCriticalextensionsChoiceNothing = iota
	UecapabilityinformationCriticalextensionsChoiceC1
	UecapabilityinformationCriticalextensionsChoiceCriticalextensionsfuture
)

type UecapabilityinformationCriticalextensions struct {
	Choice                   uint64
	C1                       *UecapabilityinformationCriticalextensionsC1
	Criticalextensionsfuture *UecapabilityinformationCriticalextensionsCriticalextensionsfuture
}
