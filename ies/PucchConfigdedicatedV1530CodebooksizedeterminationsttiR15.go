package ies

import "rrc/utils"

// PUCCH-ConfigDedicated-v1530-codebooksizeDeterminationSTTI-r15 ::= ENUMERATED
type PucchConfigdedicatedV1530CodebooksizedeterminationsttiR15 struct {
	Value utils.ENUMERATED
}

const (
	PucchConfigdedicatedV1530CodebooksizedeterminationsttiR15EnumeratedNothing = iota
	PucchConfigdedicatedV1530CodebooksizedeterminationsttiR15EnumeratedDai
	PucchConfigdedicatedV1530CodebooksizedeterminationsttiR15EnumeratedCc
)
