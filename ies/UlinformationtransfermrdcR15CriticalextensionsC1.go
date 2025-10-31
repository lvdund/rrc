package ies

// ULInformationTransferMRDC-r15-criticalExtensions-c1 ::= CHOICE
const (
	UlinformationtransfermrdcR15CriticalextensionsC1ChoiceNothing = iota
	UlinformationtransfermrdcR15CriticalextensionsC1ChoiceUlinformationtransfermrdcR15
	UlinformationtransfermrdcR15CriticalextensionsC1ChoiceSpare3
	UlinformationtransfermrdcR15CriticalextensionsC1ChoiceSpare2
	UlinformationtransfermrdcR15CriticalextensionsC1ChoiceSpare1
)

type UlinformationtransfermrdcR15CriticalextensionsC1 struct {
	Choice                       uint64
	UlinformationtransfermrdcR15 *UlinformationtransfermrdcR15
	Spare3                       *struct{}
	Spare2                       *struct{}
	Spare1                       *struct{}
}
