package ies

import "rrc/utils"

// PDCP-Parameters-ehc-r16 ::= ENUMERATED
type PdcpParametersEhcR16 struct {
	Value utils.ENUMERATED
}

const (
	PdcpParametersEhcR16EnumeratedNothing = iota
	PdcpParametersEhcR16EnumeratedSupported
)
