package ies

import "rrc/utils"

// PDCP-Parameters-v1310-pdcp-SN-Extension-18bits-r13 ::= ENUMERATED
type PdcpParametersV1310PdcpSnExtension18bitsR13 struct {
	Value utils.ENUMERATED
}

const (
	PdcpParametersV1310PdcpSnExtension18bitsR13EnumeratedNothing = iota
	PdcpParametersV1310PdcpSnExtension18bitsR13EnumeratedSupported
)
