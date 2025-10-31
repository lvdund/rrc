package ies

import "rrc/utils"

// GWUS-Paging-ProbThresh-NB-r16 ::= ENUMERATED
type GwusPagingProbthreshNbR16 struct {
	Value utils.ENUMERATED
}

const (
	GwusPagingProbthreshNbR16EnumeratedNothing = iota
	GwusPagingProbthreshNbR16EnumeratedP20
	GwusPagingProbthreshNbR16EnumeratedP30
	GwusPagingProbthreshNbR16EnumeratedP40
	GwusPagingProbthreshNbR16EnumeratedP50
	GwusPagingProbthreshNbR16EnumeratedP60
	GwusPagingProbthreshNbR16EnumeratedP70
	GwusPagingProbthreshNbR16EnumeratedP80
	GwusPagingProbthreshNbR16EnumeratedP90
)
