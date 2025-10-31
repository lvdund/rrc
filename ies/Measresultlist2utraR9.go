package ies

// MeasResultList2UTRA-r9 ::= SEQUENCE OF MeasResult2UTRA-r9
// SIZE (1..maxFreq)
type Measresultlist2utraR9 struct {
	Value []Measresult2utraR9 `lb:1,ub:maxFreq`
}
