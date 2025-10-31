package ies

// MeasResultList2EUTRA-r9 ::= SEQUENCE OF MeasResult2EUTRA-r9
// SIZE (1..maxFreq)
type Measresultlist2eutraR9 struct {
	Value []Measresult2eutraR9 `lb:1,ub:maxFreq`
}
