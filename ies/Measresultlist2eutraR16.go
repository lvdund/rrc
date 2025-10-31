package ies

// MeasResultList2EUTRA-r16 ::= SEQUENCE OF MeasResult2EUTRA-r16
// SIZE (1..maxFreq)
type Measresultlist2eutraR16 struct {
	Value []Measresult2eutraR16 `lb:1,ub:maxFreq`
}
