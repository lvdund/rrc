package ies

// MeasResultList2NR-r16 ::= SEQUENCE OF MeasResult2NR-r16
// SIZE (1..maxFreq)
type Measresultlist2nrR16 struct {
	Value []Measresult2nrR16 `lb:1,ub:maxFreq`
}
