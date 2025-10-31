package ies

// DLInformationTransferMRDC-r16-criticalExtensions ::= CHOICE
const (
	DlinformationtransfermrdcR16CriticalextensionsChoiceNothing = iota
	DlinformationtransfermrdcR16CriticalextensionsChoiceC1
	DlinformationtransfermrdcR16CriticalextensionsChoiceCriticalextensionsfuture
)

type DlinformationtransfermrdcR16Criticalextensions struct {
	Choice                   uint64
	C1                       *DlinformationtransfermrdcR16CriticalextensionsC1
	Criticalextensionsfuture *DlinformationtransfermrdcR16CriticalextensionsCriticalextensionsfuture
}
