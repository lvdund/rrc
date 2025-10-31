package ies

// MeasResultFreqListNR-r16 ::= SEQUENCE OF MeasResultFreqFailNR-r15
// SIZE (1..maxFreq-1-r16)
type MeasresultfreqlistnrR16 struct {
	Value []MeasresultfreqfailnrR15 `lb:1,ub:maxFreq1R16`
}
