package ies

import "rrc/utils"

// PDCP-Parameters-udc-r17-standardDictionary-r17 ::= ENUMERATED
type PdcpParametersUdcR17StandarddictionaryR17 struct {
	Value utils.ENUMERATED
}

const (
	PdcpParametersUdcR17StandarddictionaryR17EnumeratedNothing = iota
	PdcpParametersUdcR17StandarddictionaryR17EnumeratedSupported
)
