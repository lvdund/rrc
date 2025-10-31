package ies

import "rrc/utils"

// PDCP-Parameters-continueROHC-Context ::= ENUMERATED
type PdcpParametersContinuerohcContext struct {
	Value utils.ENUMERATED
}

const (
	PdcpParametersContinuerohcContextEnumeratedNothing = iota
	PdcpParametersContinuerohcContextEnumeratedSupported
)
