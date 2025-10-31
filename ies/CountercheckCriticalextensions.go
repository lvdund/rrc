package ies

// CounterCheck-criticalExtensions ::= CHOICE
const (
	CountercheckCriticalextensionsChoiceNothing = iota
	CountercheckCriticalextensionsChoiceC1
	CountercheckCriticalextensionsChoiceCriticalextensionsfuture
)

type CountercheckCriticalextensions struct {
	Choice                   uint64
	C1                       *CountercheckCriticalextensionsC1
	Criticalextensionsfuture *CountercheckCriticalextensionsCriticalextensionsfuture
}
