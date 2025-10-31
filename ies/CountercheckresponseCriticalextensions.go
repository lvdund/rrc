package ies

// CounterCheckResponse-criticalExtensions ::= CHOICE
const (
	CountercheckresponseCriticalextensionsChoiceNothing = iota
	CountercheckresponseCriticalextensionsChoiceCountercheckresponse
	CountercheckresponseCriticalextensionsChoiceCriticalextensionsfuture
)

type CountercheckresponseCriticalextensions struct {
	Choice                   uint64
	Countercheckresponse     *Countercheckresponse
	Criticalextensionsfuture *CountercheckresponseCriticalextensionsCriticalextensionsfuture
}
