package ies

import "rrc/utils"

// RRCConnectionResumeRequest-5GC-r15-IEs ::= SEQUENCE
type Rrcconnectionresumerequest5gcR15 struct {
	ResumeidentityR15  Rrcconnectionresumerequest5gcR15IesResumeidentityR15
	ShortresumemacIR15 utils.BITSTRING `lb:16,ub:16`
	ResumecauseR15     ResumecauseR15
	Spare              utils.BITSTRING `lb:1,ub:1`
}
