package ies

// UEInformationRequest-r16-criticalExtensions ::= CHOICE
const (
	UeinformationrequestR16CriticalextensionsChoiceNothing = iota
	UeinformationrequestR16CriticalextensionsChoiceUeinformationrequestR16
	UeinformationrequestR16CriticalextensionsChoiceCriticalextensionsfuture
)

type UeinformationrequestR16Criticalextensions struct {
	Choice                   uint64
	UeinformationrequestR16  *UeinformationrequestR16
	Criticalextensionsfuture *UeinformationrequestR16CriticalextensionsCriticalextensionsfuture
}
