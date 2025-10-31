package ies

// InterRAT-Parameters ::= SEQUENCE
// Extensible
type InterratParameters struct {
	Eutra      *EutraParameters
	UtraFddR16 *UtraFddParametersR16 `ext`
}
