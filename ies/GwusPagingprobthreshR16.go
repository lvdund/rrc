package ies

import "rrc/utils"

// GWUS-PagingProbThresh-r16 ::= ENUMERATED
type GwusPagingprobthreshR16 struct {
	Value utils.ENUMERATED
}

const (
	GwusPagingprobthreshR16EnumeratedNothing = iota
	GwusPagingprobthreshR16EnumeratedP20
	GwusPagingprobthreshR16EnumeratedP30
	GwusPagingprobthreshR16EnumeratedP40
	GwusPagingprobthreshR16EnumeratedP50
	GwusPagingprobthreshR16EnumeratedP60
	GwusPagingprobthreshR16EnumeratedP70
	GwusPagingprobthreshR16EnumeratedP80
	GwusPagingprobthreshR16EnumeratedP90
)
