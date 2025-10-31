package ies

import "rrc/utils"

// PDCP-Parameters-continueEHC-Context-r16 ::= ENUMERATED
type PdcpParametersContinueehcContextR16 struct {
	Value utils.ENUMERATED
}

const (
	PdcpParametersContinueehcContextR16EnumeratedNothing = iota
	PdcpParametersContinueehcContextR16EnumeratedSupported
)
