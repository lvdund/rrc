package ies

// DLInformationTransfer-NB-criticalExtensions-c1 ::= CHOICE
const (
	DlinformationtransferNbCriticalextensionsC1ChoiceNothing = iota
	DlinformationtransferNbCriticalextensionsC1ChoiceDlinformationtransferR13
	DlinformationtransferNbCriticalextensionsC1ChoiceSpare1
)

type DlinformationtransferNbCriticalextensionsC1 struct {
	Choice                   uint64
	DlinformationtransferR13 *DlinformationtransferNbR13
	Spare1                   *struct{}
}
