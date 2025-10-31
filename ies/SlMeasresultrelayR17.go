package ies

// SL-MeasResultRelay-r17 ::= SEQUENCE
// Extensible
type SlMeasresultrelayR17 struct {
	CellidentityR17      Cellaccessrelatedinfo
	SlRelayueIdentityR17 SlSourceidentityR17
	SlMeasresultR17      SlMeasresultR16
}
