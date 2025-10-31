package ies

import "rrc/utils"

// RRCConnectionResumeRequest-r13-IEs ::= SEQUENCE
type RrcconnectionresumerequestR13 struct {
	ResumeidentityR13  RrcconnectionresumerequestR13IesResumeidentityR13
	ShortresumemacIR13 utils.BITSTRING `lb:16,ub:16`
	ResumecauseR13     Resumecause
	Spare              utils.BITSTRING `lb:1,ub:1`
}
