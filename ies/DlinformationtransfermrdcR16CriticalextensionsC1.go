package ies

// DLInformationTransferMRDC-r16-criticalExtensions-c1 ::= CHOICE
const (
	DlinformationtransfermrdcR16CriticalextensionsC1ChoiceNothing = iota
	DlinformationtransfermrdcR16CriticalextensionsC1ChoiceDlinformationtransfermrdcR16
	DlinformationtransfermrdcR16CriticalextensionsC1ChoiceSpare3
	DlinformationtransfermrdcR16CriticalextensionsC1ChoiceSpare2
	DlinformationtransfermrdcR16CriticalextensionsC1ChoiceSpare1
)

type DlinformationtransfermrdcR16CriticalextensionsC1 struct {
	Choice                       uint64
	DlinformationtransfermrdcR16 *DlinformationtransfermrdcR16
	Spare3                       *struct{}
	Spare2                       *struct{}
	Spare1                       *struct{}
}
