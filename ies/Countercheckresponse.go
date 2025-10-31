package ies

// CounterCheckResponse ::= SEQUENCE
type Countercheckresponse struct {
	RrcTransactionidentifier RrcTransactionidentifier
	Criticalextensions       CountercheckresponseCriticalextensions
}
