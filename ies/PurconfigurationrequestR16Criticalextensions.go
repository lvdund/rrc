package ies

// PURConfigurationRequest-r16-criticalExtensions ::= CHOICE
const (
	PurconfigurationrequestR16CriticalextensionsChoiceNothing = iota
	PurconfigurationrequestR16CriticalextensionsChoicePurconfigurationrequest
	PurconfigurationrequestR16CriticalextensionsChoiceCriticalextensionsfuture
)

type PurconfigurationrequestR16Criticalextensions struct {
	Choice                   uint64
	Purconfigurationrequest  *PurconfigurationrequestR16
	Criticalextensionsfuture *PurconfigurationrequestR16CriticalextensionsCriticalextensionsfuture
}
