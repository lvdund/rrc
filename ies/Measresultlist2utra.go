package ies

// MeasResultList2UTRA ::= SEQUENCE OF MeasResult2UTRA-FDD-r16
// SIZE (1..maxFreq)
type Measresultlist2utra struct {
	Value []Measresult2utraFddR16 `lb:1,ub:maxFreq`
}
