package ies

// UEInformationRequest-r9-criticalExtensions ::= CHOICE
const (
	UeinformationrequestR9CriticalextensionsChoiceNothing = iota
	UeinformationrequestR9CriticalextensionsChoiceC1
	UeinformationrequestR9CriticalextensionsChoiceCriticalextensionsfuture
)

type UeinformationrequestR9Criticalextensions struct {
	Choice                   uint64
	C1                       *UeinformationrequestR9CriticalextensionsC1
	Criticalextensionsfuture *UeinformationrequestR9CriticalextensionsCriticalextensionsfuture
}
