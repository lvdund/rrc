package ies

// ULInformationTransferMRDC-r15-criticalExtensions ::= CHOICE
const (
	UlinformationtransfermrdcR15CriticalextensionsChoiceNothing = iota
	UlinformationtransfermrdcR15CriticalextensionsChoiceC1
	UlinformationtransfermrdcR15CriticalextensionsChoiceCriticalextensionsfuture
)

type UlinformationtransfermrdcR15Criticalextensions struct {
	Choice                   uint64
	C1                       *UlinformationtransfermrdcR15CriticalextensionsC1
	Criticalextensionsfuture *UlinformationtransfermrdcR15CriticalextensionsCriticalextensionsfuture
}
