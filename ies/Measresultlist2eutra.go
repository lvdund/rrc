package ies

// MeasResultList2EUTRA ::= SEQUENCE OF MeasResult2EUTRA-r16
// SIZE (1..maxFreq)
type Measresultlist2eutra struct {
	Value []Measresult2eutraR16 `lb:1,ub:maxFreq`
}
