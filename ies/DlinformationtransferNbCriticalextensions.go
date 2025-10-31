package ies

// DLInformationTransfer-NB-criticalExtensions ::= CHOICE
const (
	DlinformationtransferNbCriticalextensionsChoiceNothing = iota
	DlinformationtransferNbCriticalextensionsChoiceC1
	DlinformationtransferNbCriticalextensionsChoiceCriticalextensionsfuture
)

type DlinformationtransferNbCriticalextensions struct {
	Choice                   uint64
	C1                       *DlinformationtransferNbCriticalextensionsC1
	Criticalextensionsfuture *DlinformationtransferNbCriticalextensionsCriticalextensionsfuture
}
