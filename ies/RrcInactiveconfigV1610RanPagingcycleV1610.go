package ies

import "rrc/utils"

// RRC-InactiveConfig-v1610-ran-PagingCycle-v1610 ::= ENUMERATED
type RrcInactiveconfigV1610RanPagingcycleV1610 struct {
	Value utils.ENUMERATED
}

const (
	RrcInactiveconfigV1610RanPagingcycleV1610EnumeratedNothing = iota
	RrcInactiveconfigV1610RanPagingcycleV1610EnumeratedRf512
	RrcInactiveconfigV1610RanPagingcycleV1610EnumeratedRf1024
)
