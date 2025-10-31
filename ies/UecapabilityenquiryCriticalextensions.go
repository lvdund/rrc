package ies

// UECapabilityEnquiry-criticalExtensions ::= CHOICE
const (
	UecapabilityenquiryCriticalextensionsChoiceNothing = iota
	UecapabilityenquiryCriticalextensionsChoiceUecapabilityenquiry
	UecapabilityenquiryCriticalextensionsChoiceCriticalextensionsfuture
)

type UecapabilityenquiryCriticalextensions struct {
	Choice                   uint64
	Uecapabilityenquiry      *Uecapabilityenquiry
	Criticalextensionsfuture *UecapabilityenquiryCriticalextensionsCriticalextensionsfuture
}
