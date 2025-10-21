package ies

import "rrc/utils"

// GWUS-PagingProbThresh-r16 ::= ENUMERATED
type GwusPagingprobthreshR16 struct {
	Value utils.ENUMERATED
}

const (
	GwusPagingprobthreshR16P20 = 0
	GwusPagingprobthreshR16P30 = 1
	GwusPagingprobthreshR16P40 = 2
	GwusPagingprobthreshR16P50 = 3
	GwusPagingprobthreshR16P60 = 4
	GwusPagingprobthreshR16P70 = 5
	GwusPagingprobthreshR16P80 = 6
	GwusPagingprobthreshR16P90 = 7
)
