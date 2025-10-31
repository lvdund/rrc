package ies

import "rrc/utils"

// PDCP-Parameters-outOfOrderDelivery ::= ENUMERATED
type PdcpParametersOutoforderdelivery struct {
	Value utils.ENUMERATED
}

const (
	PdcpParametersOutoforderdeliveryEnumeratedNothing = iota
	PdcpParametersOutoforderdeliveryEnumeratedSupported
)
