package ies

// ULHandoverPreparationTransfer-criticalExtensions-c1 ::= CHOICE
const (
	UlhandoverpreparationtransferCriticalextensionsC1ChoiceNothing = iota
	UlhandoverpreparationtransferCriticalextensionsC1ChoiceUlhandoverpreparationtransferR8
	UlhandoverpreparationtransferCriticalextensionsC1ChoiceSpare3
	UlhandoverpreparationtransferCriticalextensionsC1ChoiceSpare2
	UlhandoverpreparationtransferCriticalextensionsC1ChoiceSpare1
)

type UlhandoverpreparationtransferCriticalextensionsC1 struct {
	Choice                          uint64
	UlhandoverpreparationtransferR8 *UlhandoverpreparationtransferR8
	Spare3                          *struct{}
	Spare2                          *struct{}
	Spare1                          *struct{}
}
