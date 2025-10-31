package ies

// ULInformationTransfer-criticalExtensions ::= CHOICE
const (
	UlinformationtransferCriticalextensionsChoiceNothing = iota
	UlinformationtransferCriticalextensionsChoiceUlinformationtransfer
	UlinformationtransferCriticalextensionsChoiceCriticalextensionsfuture
)

type UlinformationtransferCriticalextensions struct {
	Choice                   uint64
	Ulinformationtransfer    *Ulinformationtransfer
	Criticalextensionsfuture *UlinformationtransferCriticalextensionsCriticalextensionsfuture
}
