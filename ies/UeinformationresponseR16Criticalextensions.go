package ies

// UEInformationResponse-r16-criticalExtensions ::= CHOICE
const (
	UeinformationresponseR16CriticalextensionsChoiceNothing = iota
	UeinformationresponseR16CriticalextensionsChoiceUeinformationresponseR16
	UeinformationresponseR16CriticalextensionsChoiceCriticalextensionsfuture
)

type UeinformationresponseR16Criticalextensions struct {
	Choice                   uint64
	UeinformationresponseR16 *UeinformationresponseR16
	Criticalextensionsfuture *UeinformationresponseR16CriticalextensionsCriticalextensionsfuture
}
