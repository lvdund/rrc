package ies

// UEInformationResponse-NB-r16-criticalExtensions ::= CHOICE
const (
	UeinformationresponseNbR16CriticalextensionsChoiceNothing = iota
	UeinformationresponseNbR16CriticalextensionsChoiceUeinformationresponseR16
	UeinformationresponseNbR16CriticalextensionsChoiceCriticalextensionsfuture
)

type UeinformationresponseNbR16Criticalextensions struct {
	Choice                   uint64
	UeinformationresponseR16 *UeinformationresponseNbR16
	Criticalextensionsfuture *UeinformationresponseNbR16CriticalextensionsCriticalextensionsfuture
}
