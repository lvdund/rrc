package ies

import "rrc/utils"

// SL-PreconfigDiscRxPoolList-r13 ::= SEQUENCE OF SL-PreconfigDiscPool-r13
// SIZE (1..maxSL-DiscRxPoolPreconf-r13)
type SlPreconfigdiscrxpoollistR13 struct {
	Value utils.Sequence[SlPreconfigdiscpoolR13]
}
