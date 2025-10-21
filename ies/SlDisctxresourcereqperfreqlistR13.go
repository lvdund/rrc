package ies

import "rrc/utils"

// SL-DiscTxResourceReqPerFreqList-r13 ::= SEQUENCE OF SL-DiscTxResourceReq-r13
// SIZE (1..maxFreq)
type SlDisctxresourcereqperfreqlistR13 struct {
	Value utils.Sequence[SlDisctxresourcereqR13]
}
