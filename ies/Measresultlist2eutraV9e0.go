package ies

// MeasResultList2EUTRA-v9e0 ::= SEQUENCE OF MeasResult2EUTRA-v9e0
// SIZE (1..maxFreq)
type Measresultlist2eutraV9e0 struct {
	Value []Measresult2eutraV9e0 `lb:1,ub:maxFreq`
}
