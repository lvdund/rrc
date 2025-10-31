package ies

import "rrc/utils"

// PDCP-Parameters-v1610-ehc-r16 ::= ENUMERATED
type PdcpParametersV1610EhcR16 struct {
	Value utils.ENUMERATED
}

const (
	PdcpParametersV1610EhcR16EnumeratedNothing = iota
	PdcpParametersV1610EhcR16EnumeratedSupported
)
