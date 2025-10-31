package ies

// UERadioAccessCapabilityInformation-criticalExtensions ::= CHOICE
const (
	UeradioaccesscapabilityinformationCriticalextensionsChoiceNothing = iota
	UeradioaccesscapabilityinformationCriticalextensionsChoiceC1
	UeradioaccesscapabilityinformationCriticalextensionsChoiceCriticalextensionsfuture
)

type UeradioaccesscapabilityinformationCriticalextensions struct {
	Choice                   uint64
	C1                       *UeradioaccesscapabilityinformationCriticalextensionsC1
	Criticalextensionsfuture *UeradioaccesscapabilityinformationCriticalextensionsCriticalextensionsfuture
}
