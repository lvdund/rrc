package ies

// SL-TxResourceReqListCommRelay-r17 ::= SEQUENCE OF SL-TxResourceReqCommRelayInfo-r17
// SIZE (1..maxNrofSL-Dest-r16)
type SlTxresourcereqlistcommrelayR17 struct {
	Value []SlTxresourcereqcommrelayinfoR17 `lb:1,ub:maxNrofSLDestR16`
}
