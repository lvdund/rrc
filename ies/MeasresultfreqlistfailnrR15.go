package ies

// MeasResultFreqListFailNR-r15 ::= SEQUENCE OF MeasResultFreqFailNR-r15
// SIZE (1..maxFreqNR-r15)
type MeasresultfreqlistfailnrR15 struct {
	Value []MeasresultfreqfailnrR15 `lb:1,ub:maxFreqNRR15`
}
