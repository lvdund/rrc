package ies

import "rrc/utils"

// RRC-InactiveConfig-r15-ran-PagingCycle-r15 ::= ENUMERATED
type RrcInactiveconfigR15RanPagingcycleR15 struct {
	Value utils.ENUMERATED
}

const (
	RrcInactiveconfigR15RanPagingcycleR15EnumeratedNothing = iota
	RrcInactiveconfigR15RanPagingcycleR15EnumeratedRf32
	RrcInactiveconfigR15RanPagingcycleR15EnumeratedRf64
	RrcInactiveconfigR15RanPagingcycleR15EnumeratedRf128
	RrcInactiveconfigR15RanPagingcycleR15EnumeratedRf256
)
