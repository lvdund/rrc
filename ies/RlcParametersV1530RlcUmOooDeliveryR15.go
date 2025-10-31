package ies

import "rrc/utils"

// RLC-Parameters-v1530-rlc-UM-Ooo-Delivery-r15 ::= ENUMERATED
type RlcParametersV1530RlcUmOooDeliveryR15 struct {
	Value utils.ENUMERATED
}

const (
	RlcParametersV1530RlcUmOooDeliveryR15EnumeratedNothing = iota
	RlcParametersV1530RlcUmOooDeliveryR15EnumeratedSupported
)
