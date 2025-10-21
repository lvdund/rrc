package ies

import "rrc/utils"

// PDCCH-CandidateReductionsLAA-UL-r14 ::= CHOICE
type PdcchCandidatereductionslaaUlR14 interface {
	isPdcchCandidatereductionslaaUlR14()
}

type PdcchCandidatereductionslaaUlR14Release struct {
	Value struct{}
}

func (*PdcchCandidatereductionslaaUlR14Release) isPdcchCandidatereductionslaaUlR14() {}

type PdcchCandidatereductionslaaUlR14Setup struct {
	Value interface{}
}

func (*PdcchCandidatereductionslaaUlR14Setup) isPdcchCandidatereductionslaaUlR14() {}
