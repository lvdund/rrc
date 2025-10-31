package ies

// SL-TxResourceReqList-r16 ::= SEQUENCE OF SL-TxResourceReq-r16
// SIZE (1..maxNrofSL-Dest-r16)
type SlTxresourcereqlistR16 struct {
	Value []SlTxresourcereqR16 `lb:1,ub:maxNrofSLDestR16`
}
