package ies

// UECapabilityInformation-criticalExtensions ::= CHOICE
const (
	UecapabilityinformationCriticalextensionsChoiceNothing = iota
	UecapabilityinformationCriticalextensionsChoiceUecapabilityinformation
	UecapabilityinformationCriticalextensionsChoiceCriticalextensionsfuture
)

type UecapabilityinformationCriticalextensions struct {
	Choice                   uint64
	Uecapabilityinformation  *Uecapabilityinformation
	Criticalextensionsfuture *UecapabilityinformationCriticalextensionsCriticalextensionsfuture
}
