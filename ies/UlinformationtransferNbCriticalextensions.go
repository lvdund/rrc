package ies

// ULInformationTransfer-NB-criticalExtensions ::= CHOICE
const (
	UlinformationtransferNbCriticalextensionsChoiceNothing = iota
	UlinformationtransferNbCriticalextensionsChoiceUlinformationtransferR13
	UlinformationtransferNbCriticalextensionsChoiceCriticalextensionsfuture
)

type UlinformationtransferNbCriticalextensions struct {
	Choice                   uint64
	UlinformationtransferR13 *UlinformationtransferNbR13
	Criticalextensionsfuture *UlinformationtransferNbCriticalextensionsCriticalextensionsfuture
}
