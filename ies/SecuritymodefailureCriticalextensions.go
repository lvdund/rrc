package ies

// SecurityModeFailure-criticalExtensions ::= CHOICE
const (
	SecuritymodefailureCriticalextensionsChoiceNothing = iota
	SecuritymodefailureCriticalextensionsChoiceSecuritymodefailure
	SecuritymodefailureCriticalextensionsChoiceCriticalextensionsfuture
)

type SecuritymodefailureCriticalextensions struct {
	Choice                   uint64
	Securitymodefailure      *Securitymodefailure
	Criticalextensionsfuture *SecuritymodefailureCriticalextensionsCriticalextensionsfuture
}
