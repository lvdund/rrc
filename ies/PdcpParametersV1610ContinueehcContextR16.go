package ies

import "rrc/utils"

// PDCP-Parameters-v1610-continueEHC-Context-r16 ::= ENUMERATED
type PdcpParametersV1610ContinueehcContextR16 struct {
	Value utils.ENUMERATED
}

const (
	PdcpParametersV1610ContinueehcContextR16EnumeratedNothing = iota
	PdcpParametersV1610ContinueehcContextR16EnumeratedSupported
)
