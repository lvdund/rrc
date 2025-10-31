package ies

import "rrc/utils"

// PDCP-Parameters-extendedDiscardTimer-r16 ::= ENUMERATED
type PdcpParametersExtendeddiscardtimerR16 struct {
	Value utils.ENUMERATED
}

const (
	PdcpParametersExtendeddiscardtimerR16EnumeratedNothing = iota
	PdcpParametersExtendeddiscardtimerR16EnumeratedSupported
)
