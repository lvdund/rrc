package ies

// SL-V2X-PreconfigFreqList-r14 ::= SEQUENCE OF SL-V2X-PreconfigFreqInfo-r14
// SIZE (1..maxFreqV2X-r14)
type SlV2xPreconfigfreqlistR14 struct {
	Value []SlV2xPreconfigfreqinfoR14 `lb:1,ub:maxFreqV2XR14`
}
