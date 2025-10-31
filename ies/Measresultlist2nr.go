package ies

// MeasResultList2NR ::= SEQUENCE OF MeasResult2NR
// SIZE (1..maxFreq)
type Measresultlist2nr struct {
	Value []Measresult2nr `lb:1,ub:maxFreq`
}
