package ies

// UECapabilityEnquirySidelink-criticalExtensions ::= CHOICE
const (
	UecapabilityenquirysidelinkCriticalextensionsChoiceNothing = iota
	UecapabilityenquirysidelinkCriticalextensionsChoiceUecapabilityenquirysidelinkR16
	UecapabilityenquirysidelinkCriticalextensionsChoiceCriticalextensionsfuture
)

type UecapabilityenquirysidelinkCriticalextensions struct {
	Choice                         uint64
	UecapabilityenquirysidelinkR16 *UecapabilityenquirysidelinkR16
	Criticalextensionsfuture       *UecapabilityenquirysidelinkCriticalextensionsCriticalextensionsfuture
}
