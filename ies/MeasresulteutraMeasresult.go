package ies

// MeasResultEUTRA-measResult ::= SEQUENCE
// Extensible
type MeasresulteutraMeasresult struct {
	Rsrpresult *RsrpRange
	Rsrqresult *RsrqRange
}
