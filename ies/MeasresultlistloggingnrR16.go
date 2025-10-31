package ies

// MeasResultListLoggingNR-r16 ::= SEQUENCE OF MeasResultLoggingNR-r16
// SIZE (1..maxCellReport)
type MeasresultlistloggingnrR16 struct {
	Value []MeasresultloggingnrR16 `lb:1,ub:maxCellReport`
}
