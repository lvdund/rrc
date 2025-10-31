package ies

// SecurityModeFailure-criticalExtensions ::= CHOICE
const (
	SecuritymodefailureCriticalextensionsChoiceNothing = iota
	SecuritymodefailureCriticalextensionsChoiceSecuritymodefailureR8
	SecuritymodefailureCriticalextensionsChoiceCriticalextensionsfuture
)

type SecuritymodefailureCriticalextensions struct {
	Choice                   uint64
	SecuritymodefailureR8    *SecuritymodefailureR8
	Criticalextensionsfuture *SecuritymodefailureCriticalextensionsCriticalextensionsfuture
}
