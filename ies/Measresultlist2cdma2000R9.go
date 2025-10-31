package ies

// MeasResultList2CDMA2000-r9 ::= SEQUENCE OF MeasResult2CDMA2000-r9
// SIZE (1..maxFreq)
type Measresultlist2cdma2000R9 struct {
	Value []Measresult2cdma2000R9 `lb:1,ub:maxFreq`
}
