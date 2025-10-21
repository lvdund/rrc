package ies

import "rrc/utils"

// PDCCH-CandidateReductionValue-r14 ::= ENUMERATED
type PdcchCandidatereductionvalueR14 struct {
	Value utils.ENUMERATED
}

const (
	PdcchCandidatereductionvalueR14N0   = 0
	PdcchCandidatereductionvalueR14N50  = 1
	PdcchCandidatereductionvalueR14N100 = 2
	PdcchCandidatereductionvalueR14N150 = 3
)
