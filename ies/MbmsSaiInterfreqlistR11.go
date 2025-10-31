package ies

// MBMS-SAI-InterFreqList-r11 ::= SEQUENCE OF MBMS-SAI-InterFreq-r11
// SIZE (1..maxFreq)
type MbmsSaiInterfreqlistR11 struct {
	Value []MbmsSaiInterfreqR11 `lb:1,ub:maxFreq`
}
