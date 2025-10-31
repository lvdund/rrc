package ies

// SecurityModeCommand-criticalExtensions ::= CHOICE
const (
	SecuritymodecommandCriticalextensionsChoiceNothing = iota
	SecuritymodecommandCriticalextensionsChoiceSecuritymodecommand
	SecuritymodecommandCriticalextensionsChoiceCriticalextensionsfuture
)

type SecuritymodecommandCriticalextensions struct {
	Choice                   uint64
	Securitymodecommand      *Securitymodecommand
	Criticalextensionsfuture *SecuritymodecommandCriticalextensionsCriticalextensionsfuture
}
