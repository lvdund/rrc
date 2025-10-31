package ies

import "rrc/utils"

// MRB-PDCP-ConfigBroadcast-r17-pdcp-SN-SizeDL-r17 ::= ENUMERATED
type MrbPdcpConfigbroadcastR17PdcpSnSizedlR17 struct {
	Value utils.ENUMERATED
}

const (
	MrbPdcpConfigbroadcastR17PdcpSnSizedlR17EnumeratedNothing = iota
	MrbPdcpConfigbroadcastR17PdcpSnSizedlR17EnumeratedLen12bits
)
