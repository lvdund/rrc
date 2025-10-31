package ies

import "rrc/utils"

// PDCP-Parameters-udc-r17-continueUDC-r17 ::= ENUMERATED
type PdcpParametersUdcR17ContinueudcR17 struct {
	Value utils.ENUMERATED
}

const (
	PdcpParametersUdcR17ContinueudcR17EnumeratedNothing = iota
	PdcpParametersUdcR17ContinueudcR17EnumeratedSupported
)
