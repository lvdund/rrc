package ies

// UEAssistanceInformation-r11-criticalExtensions ::= CHOICE
const (
	UeassistanceinformationR11CriticalextensionsChoiceNothing = iota
	UeassistanceinformationR11CriticalextensionsChoiceC1
	UeassistanceinformationR11CriticalextensionsChoiceCriticalextensionsfuture
)

type UeassistanceinformationR11Criticalextensions struct {
	Choice                   uint64
	C1                       *UeassistanceinformationR11CriticalextensionsC1
	Criticalextensionsfuture *UeassistanceinformationR11CriticalextensionsCriticalextensionsfuture
}
