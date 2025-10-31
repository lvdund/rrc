package ies

// ULInformationTransferIRAT-r16-criticalExtensions-c1 ::= CHOICE
const (
	UlinformationtransferiratR16CriticalextensionsC1ChoiceNothing = iota
	UlinformationtransferiratR16CriticalextensionsC1ChoiceUlinformationtransferiratR16
	UlinformationtransferiratR16CriticalextensionsC1ChoiceSpare3
	UlinformationtransferiratR16CriticalextensionsC1ChoiceSpare2
	UlinformationtransferiratR16CriticalextensionsC1ChoiceSpare1
)

type UlinformationtransferiratR16CriticalextensionsC1 struct {
	Choice                       uint64
	UlinformationtransferiratR16 *UlinformationtransferiratR16
	Spare3                       *struct{}
	Spare2                       *struct{}
	Spare1                       *struct{}
}
