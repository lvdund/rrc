package ies

// MeasResultFreqList ::= SEQUENCE OF MeasResult2NR
// SIZE (1..maxFreq)
type Measresultfreqlist struct {
	Value []Measresult2nr `lb:1,ub:maxFreq`
}
