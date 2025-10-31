package ies

import "rrc/utils"

// RLC-Parameters-v1530-rlc-AM-Ooo-Delivery-r15 ::= ENUMERATED
type RlcParametersV1530RlcAmOooDeliveryR15 struct {
	Value utils.ENUMERATED
}

const (
	RlcParametersV1530RlcAmOooDeliveryR15EnumeratedNothing = iota
	RlcParametersV1530RlcAmOooDeliveryR15EnumeratedSupported
)
