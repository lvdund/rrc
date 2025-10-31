package ies

// UEInformationRequest-NB-r16-criticalExtensions ::= CHOICE
const (
	UeinformationrequestNbR16CriticalextensionsChoiceNothing = iota
	UeinformationrequestNbR16CriticalextensionsChoiceUeinformationrequestR16
	UeinformationrequestNbR16CriticalextensionsChoiceCriticalextensionsfuture
)

type UeinformationrequestNbR16Criticalextensions struct {
	Choice                   uint64
	UeinformationrequestR16  *UeinformationrequestNbR16
	Criticalextensionsfuture *UeinformationrequestNbR16CriticalextensionsCriticalextensionsfuture
}
