package ies

import "rrc/utils"

// PDCP-Parameters-shortSN ::= ENUMERATED
type PdcpParametersShortsn struct {
	Value utils.ENUMERATED
}

const (
	PdcpParametersShortsnEnumeratedNothing = iota
	PdcpParametersShortsnEnumeratedSupported
)
