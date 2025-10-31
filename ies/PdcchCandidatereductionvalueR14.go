package ies

import "rrc/utils"

// PDCCH-CandidateReductionValue-r14 ::= ENUMERATED
type PdcchCandidatereductionvalueR14 struct {
	Value utils.ENUMERATED
}

const (
	PdcchCandidatereductionvalueR14EnumeratedNothing = iota
	PdcchCandidatereductionvalueR14EnumeratedN0
	PdcchCandidatereductionvalueR14EnumeratedN50
	PdcchCandidatereductionvalueR14EnumeratedN100
	PdcchCandidatereductionvalueR14EnumeratedN150
)
