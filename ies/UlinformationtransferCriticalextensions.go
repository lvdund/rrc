package ies

// ULInformationTransfer-criticalExtensions ::= CHOICE
const (
	UlinformationtransferCriticalextensionsChoiceNothing = iota
	UlinformationtransferCriticalextensionsChoiceC1
	UlinformationtransferCriticalextensionsChoiceCriticalextensionsfuture
)

type UlinformationtransferCriticalextensions struct {
	Choice                   uint64
	C1                       *UlinformationtransferCriticalextensionsC1
	Criticalextensionsfuture *UlinformationtransferCriticalextensionsCriticalextensionsfuture
}
