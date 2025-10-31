package ies

// UEAssistanceInformation-criticalExtensions ::= CHOICE
const (
	UeassistanceinformationCriticalextensionsChoiceNothing = iota
	UeassistanceinformationCriticalextensionsChoiceUeassistanceinformation
	UeassistanceinformationCriticalextensionsChoiceCriticalextensionsfuture
)

type UeassistanceinformationCriticalextensions struct {
	Choice                   uint64
	Ueassistanceinformation  *Ueassistanceinformation
	Criticalextensionsfuture *UeassistanceinformationCriticalextensionsCriticalextensionsfuture
}
