package ies

// SL-GapRequest-r13 ::= SEQUENCE OF SL-GapFreqInfo-r13
// SIZE (1..maxFreq)
type SlGaprequestR13 struct {
	Value []SlGapfreqinfoR13 `lb:1,ub:maxFreq`
}
