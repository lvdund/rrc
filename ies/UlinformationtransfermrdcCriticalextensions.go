package ies

// ULInformationTransferMRDC-criticalExtensions ::= CHOICE
const (
	UlinformationtransfermrdcCriticalextensionsChoiceNothing = iota
	UlinformationtransfermrdcCriticalextensionsChoiceC1
	UlinformationtransfermrdcCriticalextensionsChoiceCriticalextensionsfuture
)

type UlinformationtransfermrdcCriticalextensions struct {
	Choice                   uint64
	C1                       *UlinformationtransfermrdcCriticalextensionsC1
	Criticalextensionsfuture *UlinformationtransfermrdcCriticalextensionsCriticalextensionsfuture
}
