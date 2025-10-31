package ies

// ULInformationTransfer-criticalExtensions-c1 ::= CHOICE
const (
	UlinformationtransferCriticalextensionsC1ChoiceNothing = iota
	UlinformationtransferCriticalextensionsC1ChoiceUlinformationtransferR8
	UlinformationtransferCriticalextensionsC1ChoiceUlinformationtransferR16
	UlinformationtransferCriticalextensionsC1ChoiceSpare2
	UlinformationtransferCriticalextensionsC1ChoiceSpare1
)

type UlinformationtransferCriticalextensionsC1 struct {
	Choice                   uint64
	UlinformationtransferR8  *UlinformationtransferR8
	UlinformationtransferR16 *UlinformationtransferR16
	Spare2                   *struct{}
	Spare1                   *struct{}
}
