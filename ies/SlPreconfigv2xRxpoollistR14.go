package ies

import "rrc/utils"

// SL-PreconfigV2X-RxPoolList-r14 ::= SEQUENCE OF SL-V2X-PreconfigCommPool-r14
// SIZE (1..maxSL-V2X-RxPoolPreconf-r14)
type SlPreconfigv2xRxpoollistR14 struct {
	Value utils.Sequence[SlV2xPreconfigcommpoolR14]
}
