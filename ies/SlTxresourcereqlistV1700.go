package ies

// SL-TxResourceReqList-v1700 ::= SEQUENCE OF SL-TxResourceReq-v1700
// SIZE (1..maxNrofSL-Dest-r16)
type SlTxresourcereqlistV1700 struct {
	Value []SlTxresourcereqV1700 `lb:1,ub:maxNrofSLDestR16`
}
