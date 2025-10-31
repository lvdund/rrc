package ies

// CounterCheck-criticalExtensions-c1 ::= CHOICE
const (
	CountercheckCriticalextensionsC1ChoiceNothing = iota
	CountercheckCriticalextensionsC1ChoiceCountercheckR8
	CountercheckCriticalextensionsC1ChoiceSpare3
	CountercheckCriticalextensionsC1ChoiceSpare2
	CountercheckCriticalextensionsC1ChoiceSpare1
)

type CountercheckCriticalextensionsC1 struct {
	Choice         uint64
	CountercheckR8 *CountercheckR8
	Spare3         *struct{}
	Spare2         *struct{}
	Spare1         *struct{}
}
