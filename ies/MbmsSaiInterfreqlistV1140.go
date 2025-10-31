package ies

// MBMS-SAI-InterFreqList-v1140 ::= SEQUENCE OF MBMS-SAI-InterFreq-v1140
// SIZE (1..maxFreq)
type MbmsSaiInterfreqlistV1140 struct {
	Value []MbmsSaiInterfreqV1140 `lb:1,ub:maxFreq`
}
