package ies

// DLInformationTransfer-criticalExtensions ::= CHOICE
const (
	DlinformationtransferCriticalextensionsChoiceNothing = iota
	DlinformationtransferCriticalextensionsChoiceDlinformationtransfer
	DlinformationtransferCriticalextensionsChoiceCriticalextensionsfuture
)

type DlinformationtransferCriticalextensions struct {
	Choice                   uint64
	Dlinformationtransfer    *Dlinformationtransfer
	Criticalextensionsfuture *DlinformationtransferCriticalextensionsCriticalextensionsfuture
}
