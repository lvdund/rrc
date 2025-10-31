package ies

// MRB-PDCP-ConfigBroadcast-r17 ::= SEQUENCE
type MrbPdcpConfigbroadcastR17 struct {
	PdcpSnSizedlR17      *MrbPdcpConfigbroadcastR17PdcpSnSizedlR17
	HeadercompressionR17 MrbPdcpConfigbroadcastR17HeadercompressionR17
	TReorderingR17       *MrbPdcpConfigbroadcastR17TReorderingR17
}
