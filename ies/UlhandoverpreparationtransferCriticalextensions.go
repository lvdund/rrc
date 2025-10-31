package ies

// ULHandoverPreparationTransfer-criticalExtensions ::= CHOICE
const (
	UlhandoverpreparationtransferCriticalextensionsChoiceNothing = iota
	UlhandoverpreparationtransferCriticalextensionsChoiceC1
	UlhandoverpreparationtransferCriticalextensionsChoiceCriticalextensionsfuture
)

type UlhandoverpreparationtransferCriticalextensions struct {
	Choice                   uint64
	C1                       *UlhandoverpreparationtransferCriticalextensionsC1
	Criticalextensionsfuture *UlhandoverpreparationtransferCriticalextensionsCriticalextensionsfuture
}
