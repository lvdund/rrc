package ies

// SecurityModeCommand-criticalExtensions ::= CHOICE
const (
	SecuritymodecommandCriticalextensionsChoiceNothing = iota
	SecuritymodecommandCriticalextensionsChoiceC1
	SecuritymodecommandCriticalextensionsChoiceCriticalextensionsfuture
)

type SecuritymodecommandCriticalextensions struct {
	Choice                   uint64
	C1                       *SecuritymodecommandCriticalextensionsC1
	Criticalextensionsfuture *SecuritymodecommandCriticalextensionsCriticalextensionsfuture
}
