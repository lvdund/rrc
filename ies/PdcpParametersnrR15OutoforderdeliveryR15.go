package ies

import "rrc/utils"

// PDCP-ParametersNR-r15-outOfOrderDelivery-r15 ::= ENUMERATED
type PdcpParametersnrR15OutoforderdeliveryR15 struct {
	Value utils.ENUMERATED
}

const (
	PdcpParametersnrR15OutoforderdeliveryR15EnumeratedNothing = iota
	PdcpParametersnrR15OutoforderdeliveryR15EnumeratedSupported
)
