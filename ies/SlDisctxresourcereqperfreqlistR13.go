package ies

// SL-DiscTxResourceReqPerFreqList-r13 ::= SEQUENCE OF SL-DiscTxResourceReq-r13
// SIZE (1..maxFreq)
type SlDisctxresourcereqperfreqlistR13 struct {
	Value []SlDisctxresourcereqR13 `lb:1,ub:maxFreq`
}
