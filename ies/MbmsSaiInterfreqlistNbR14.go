package ies

// MBMS-SAI-InterFreqList-NB-r14 ::= SEQUENCE OF MBMS-SAI-InterFreq-NB-r14
// SIZE (1..maxFreq)
type MbmsSaiInterfreqlistNbR14 struct {
	Value []MbmsSaiInterfreqNbR14 `lb:1,ub:maxFreq`
}
