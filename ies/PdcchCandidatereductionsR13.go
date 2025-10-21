package ies

import "rrc/utils"

// PDCCH-CandidateReductions-r13 ::= CHOICE
type PdcchCandidatereductionsR13 interface {
	isPdcchCandidatereductionsR13()
}

type PdcchCandidatereductionsR13Release struct {
	Value struct{}
}

func (*PdcchCandidatereductionsR13Release) isPdcchCandidatereductionsR13() {}

type PdcchCandidatereductionsR13Setup struct {
	Value interface{}
}

func (*PdcchCandidatereductionsR13Setup) isPdcchCandidatereductionsR13() {}
