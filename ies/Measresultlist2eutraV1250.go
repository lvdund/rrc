package ies

// MeasResultList2EUTRA-v1250 ::= SEQUENCE OF MeasResult2EUTRA-v1250
// SIZE (1..maxFreq)
type Measresultlist2eutraV1250 struct {
	Value []Measresult2eutraV1250 `lb:1,ub:maxFreq`
}
