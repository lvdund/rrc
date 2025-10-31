package ies

import "rrc/utils"

// PagingCycle ::= ENUMERATED
type Pagingcycle struct {
	Value utils.ENUMERATED
}

const (
	PagingcycleEnumeratedNothing = iota
	PagingcycleEnumeratedRf32
	PagingcycleEnumeratedRf64
	PagingcycleEnumeratedRf128
	PagingcycleEnumeratedRf256
)
