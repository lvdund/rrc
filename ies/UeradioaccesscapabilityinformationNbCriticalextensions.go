package ies

// UERadioAccessCapabilityInformation-NB-criticalExtensions ::= CHOICE
const (
	UeradioaccesscapabilityinformationNbCriticalextensionsChoiceNothing = iota
	UeradioaccesscapabilityinformationNbCriticalextensionsChoiceC1
	UeradioaccesscapabilityinformationNbCriticalextensionsChoiceCriticalextensionsfuture
)

type UeradioaccesscapabilityinformationNbCriticalextensions struct {
	Choice                   uint64
	C1                       *UeradioaccesscapabilityinformationNbCriticalextensionsC1
	Criticalextensionsfuture *UeradioaccesscapabilityinformationNbCriticalextensionsCriticalextensionsfuture
}
