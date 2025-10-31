package ies

import "rrc/utils"

// PDCCH-CandidateReductionValue-r13 ::= ENUMERATED
type PdcchCandidatereductionvalueR13 struct {
	Value utils.ENUMERATED
}

const (
	PdcchCandidatereductionvalueR13EnumeratedNothing = iota
	PdcchCandidatereductionvalueR13EnumeratedN0
	PdcchCandidatereductionvalueR13EnumeratedN33
	PdcchCandidatereductionvalueR13EnumeratedN66
	PdcchCandidatereductionvalueR13EnumeratedN100
)
