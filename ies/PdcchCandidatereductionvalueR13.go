package ies

import "rrc/utils"

// PDCCH-CandidateReductionValue-r13 ::= ENUMERATED
type PdcchCandidatereductionvalueR13 struct {
	Value utils.ENUMERATED
}

const (
	PdcchCandidatereductionvalueR13N0   = 0
	PdcchCandidatereductionvalueR13N33  = 1
	PdcchCandidatereductionvalueR13N66  = 2
	PdcchCandidatereductionvalueR13N100 = 3
)
