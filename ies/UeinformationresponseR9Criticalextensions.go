package ies

// UEInformationResponse-r9-criticalExtensions ::= CHOICE
const (
	UeinformationresponseR9CriticalextensionsChoiceNothing = iota
	UeinformationresponseR9CriticalextensionsChoiceC1
	UeinformationresponseR9CriticalextensionsChoiceCriticalextensionsfuture
)

type UeinformationresponseR9Criticalextensions struct {
	Choice                   uint64
	C1                       *UeinformationresponseR9CriticalextensionsC1
	Criticalextensionsfuture *UeinformationresponseR9CriticalextensionsCriticalextensionsfuture
}
