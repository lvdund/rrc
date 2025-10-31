package ies

// PURConfigurationRequest-NB-r16-criticalExtensions ::= CHOICE
const (
	PurconfigurationrequestNbR16CriticalextensionsChoiceNothing = iota
	PurconfigurationrequestNbR16CriticalextensionsChoicePurconfigurationrequestR16
	PurconfigurationrequestNbR16CriticalextensionsChoiceCriticalextensionsfuture
)

type PurconfigurationrequestNbR16Criticalextensions struct {
	Choice                     uint64
	PurconfigurationrequestR16 *PurconfigurationrequestNbR16
	Criticalextensionsfuture   *PurconfigurationrequestNbR16CriticalextensionsCriticalextensionsfuture
}
