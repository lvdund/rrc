package ies

// SecurityModeComplete-criticalExtensions ::= CHOICE
const (
	SecuritymodecompleteCriticalextensionsChoiceNothing = iota
	SecuritymodecompleteCriticalextensionsChoiceSecuritymodecompleteR8
	SecuritymodecompleteCriticalextensionsChoiceCriticalextensionsfuture
)

type SecuritymodecompleteCriticalextensions struct {
	Choice                   uint64
	SecuritymodecompleteR8   *SecuritymodecompleteR8
	Criticalextensionsfuture *SecuritymodecompleteCriticalextensionsCriticalextensionsfuture
}
