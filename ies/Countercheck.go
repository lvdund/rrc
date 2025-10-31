package ies

// CounterCheck ::= SEQUENCE
type Countercheck struct {
	RrcTransactionidentifier RrcTransactionidentifier
	Criticalextensions       CountercheckCriticalextensions
}
