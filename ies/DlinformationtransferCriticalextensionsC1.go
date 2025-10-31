package ies

// DLInformationTransfer-criticalExtensions-c1 ::= CHOICE
const (
	DlinformationtransferCriticalextensionsC1ChoiceNothing = iota
	DlinformationtransferCriticalextensionsC1ChoiceDlinformationtransferR8
	DlinformationtransferCriticalextensionsC1ChoiceDlinformationtransferR15
	DlinformationtransferCriticalextensionsC1ChoiceSpare2
	DlinformationtransferCriticalextensionsC1ChoiceSpare1
)

type DlinformationtransferCriticalextensionsC1 struct {
	Choice                   uint64
	DlinformationtransferR8  *DlinformationtransferR8
	DlinformationtransferR15 *DlinformationtransferR15
	Spare2                   *struct{}
	Spare1                   *struct{}
}
