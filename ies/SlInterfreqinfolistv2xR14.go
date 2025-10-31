package ies

// SL-InterFreqInfoListV2X-r14 ::= SEQUENCE OF SL-InterFreqInfoV2X-r14
// SIZE (0..maxFreqV2X-1-r14)
type SlInterfreqinfolistv2xR14 struct {
	Value []SlInterfreqinfov2xR14 `lb:0,ub:maxFreqV2X1R14`
}
