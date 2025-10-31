package ies

// UECapabilityEnquiry-criticalExtensions ::= CHOICE
const (
	UecapabilityenquiryCriticalextensionsChoiceNothing = iota
	UecapabilityenquiryCriticalextensionsChoiceC1
	UecapabilityenquiryCriticalextensionsChoiceCriticalextensionsfuture
)

type UecapabilityenquiryCriticalextensions struct {
	Choice                   uint64
	C1                       *UecapabilityenquiryCriticalextensionsC1
	Criticalextensionsfuture *UecapabilityenquiryCriticalextensionsCriticalextensionsfuture
}
