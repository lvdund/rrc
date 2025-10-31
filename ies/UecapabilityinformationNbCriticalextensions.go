package ies

// UECapabilityInformation-NB-criticalExtensions ::= CHOICE
const (
	UecapabilityinformationNbCriticalextensionsChoiceNothing = iota
	UecapabilityinformationNbCriticalextensionsChoiceUecapabilityinformationR13
	UecapabilityinformationNbCriticalextensionsChoiceCriticalextensionsfuture
)

type UecapabilityinformationNbCriticalextensions struct {
	Choice                     uint64
	UecapabilityinformationR13 *UecapabilityinformationNbR13
	Criticalextensionsfuture   *UecapabilityinformationNbCriticalextensionsCriticalextensionsfuture
}
