package ies

// SecurityModeComplete-criticalExtensions ::= CHOICE
const (
	SecuritymodecompleteCriticalextensionsChoiceNothing = iota
	SecuritymodecompleteCriticalextensionsChoiceSecuritymodecomplete
	SecuritymodecompleteCriticalextensionsChoiceCriticalextensionsfuture
)

type SecuritymodecompleteCriticalextensions struct {
	Choice                   uint64
	Securitymodecomplete     *Securitymodecomplete
	Criticalextensionsfuture *SecuritymodecompleteCriticalextensionsCriticalextensionsfuture
}
