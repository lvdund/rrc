package ies

import "rrc/utils"

// PUCCH-ConfigDedicated-r13-codebooksizeDetermination-r13 ::= ENUMERATED
type PucchConfigdedicatedR13CodebooksizedeterminationR13 struct {
	Value utils.ENUMERATED
}

const (
	PucchConfigdedicatedR13CodebooksizedeterminationR13EnumeratedNothing = iota
	PucchConfigdedicatedR13CodebooksizedeterminationR13EnumeratedDai
	PucchConfigdedicatedR13CodebooksizedeterminationR13EnumeratedCc
)
