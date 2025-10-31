package ies

// MeasResultListLogging2NR-r16 ::= SEQUENCE OF MeasResultLogging2NR-r16
// SIZE (1..maxFreq)
type Measresultlistlogging2nrR16 struct {
	Value []Measresultlogging2nrR16 `lb:1,ub:maxFreq`
}
