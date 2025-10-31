package ies

// HandoverCommand-criticalExtensions ::= CHOICE
const (
	HandovercommandCriticalextensionsChoiceNothing = iota
	HandovercommandCriticalextensionsChoiceC1
	HandovercommandCriticalextensionsChoiceCriticalextensionsfuture
)

type HandovercommandCriticalextensions struct {
	Choice                   uint64
	C1                       *HandovercommandCriticalextensionsC1
	Criticalextensionsfuture *HandovercommandCriticalextensionsCriticalextensionsfuture
}
