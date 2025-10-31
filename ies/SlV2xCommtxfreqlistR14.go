package ies

// SL-V2X-CommTxFreqList-r14 ::= SEQUENCE OF SL-V2X-CommTxResourceReq-r14
// SIZE (1..maxFreqV2X-r14)
type SlV2xCommtxfreqlistR14 struct {
	Value []SlV2xCommtxresourcereqR14 `lb:1,ub:maxFreqV2XR14`
}
