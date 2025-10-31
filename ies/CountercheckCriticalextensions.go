package ies

// CounterCheck-criticalExtensions ::= CHOICE
const (
	CountercheckCriticalextensionsChoiceNothing = iota
	CountercheckCriticalextensionsChoiceCountercheck
	CountercheckCriticalextensionsChoiceCriticalextensionsfuture
)

type CountercheckCriticalextensions struct {
	Choice                   uint64
	Countercheck             *Countercheck
	Criticalextensionsfuture *CountercheckCriticalextensionsCriticalextensionsfuture
}
