package ies

// MeasResultFreqListFailMRDC ::= SEQUENCE OF MeasResult2EUTRA
// SIZE (1.. maxFreq)
type Measresultfreqlistfailmrdc struct {
	Value []Measresult2eutra `lb:1,ub:maxFreq`
}
