package ies

// SL-TxResourceReqListDisc-r17 ::= SEQUENCE OF SL-TxResourceReqDisc-r17
// SIZE (1..maxNrofSL-Dest-r16)
type SlTxresourcereqlistdiscR17 struct {
	Value []SlTxresourcereqdiscR17 `lb:1,ub:maxNrofSLDestR16`
}
