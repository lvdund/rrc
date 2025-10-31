package ies

import "rrc/utils"

// PDCP-ParametersSidelink-r16-outOfOrderDeliverySidelink-r16 ::= ENUMERATED
type PdcpParameterssidelinkR16OutoforderdeliverysidelinkR16 struct {
	Value utils.ENUMERATED
}

const (
	PdcpParameterssidelinkR16OutoforderdeliverysidelinkR16EnumeratedNothing = iota
	PdcpParameterssidelinkR16OutoforderdeliverysidelinkR16EnumeratedSupported
)
