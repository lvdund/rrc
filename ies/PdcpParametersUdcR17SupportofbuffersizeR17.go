package ies

import "rrc/utils"

// PDCP-Parameters-udc-r17-supportOfBufferSize-r17 ::= ENUMERATED
type PdcpParametersUdcR17SupportofbuffersizeR17 struct {
	Value utils.ENUMERATED
}

const (
	PdcpParametersUdcR17SupportofbuffersizeR17EnumeratedNothing = iota
	PdcpParametersUdcR17SupportofbuffersizeR17EnumeratedKbyte4
	PdcpParametersUdcR17SupportofbuffersizeR17EnumeratedKbyte8
)
