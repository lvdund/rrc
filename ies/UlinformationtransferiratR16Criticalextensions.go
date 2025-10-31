package ies

// ULInformationTransferIRAT-r16-criticalExtensions ::= CHOICE
const (
	UlinformationtransferiratR16CriticalextensionsChoiceNothing = iota
	UlinformationtransferiratR16CriticalextensionsChoiceC1
	UlinformationtransferiratR16CriticalextensionsChoiceCriticalextensionsfuture
)

type UlinformationtransferiratR16Criticalextensions struct {
	Choice                   uint64
	C1                       *UlinformationtransferiratR16CriticalextensionsC1
	Criticalextensionsfuture *UlinformationtransferiratR16CriticalextensionsCriticalextensionsfuture
}
