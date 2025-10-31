package ies

// CounterCheckResponse-criticalExtensions ::= CHOICE
const (
	CountercheckresponseCriticalextensionsChoiceNothing = iota
	CountercheckresponseCriticalextensionsChoiceCountercheckresponseR8
	CountercheckresponseCriticalextensionsChoiceCriticalextensionsfuture
)

type CountercheckresponseCriticalextensions struct {
	Choice                   uint64
	CountercheckresponseR8   *CountercheckresponseR8
	Criticalextensionsfuture *CountercheckresponseCriticalextensionsCriticalextensionsfuture
}
