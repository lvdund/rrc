package ies

// DLInformationTransfer-criticalExtensions ::= CHOICE
const (
	DlinformationtransferCriticalextensionsChoiceNothing = iota
	DlinformationtransferCriticalextensionsChoiceC1
	DlinformationtransferCriticalextensionsChoiceCriticalextensionsfuture
)

type DlinformationtransferCriticalextensions struct {
	Choice                   uint64
	C1                       *DlinformationtransferCriticalextensionsC1
	Criticalextensionsfuture *DlinformationtransferCriticalextensionsCriticalextensionsfuture
}
