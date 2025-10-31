package ies

import "rrc/utils"

// RRCResumeRequest1-IEs ::= SEQUENCE
type Rrcresumerequest1 struct {
	Resumeidentity IRntiValue
	ResumemacI     utils.BITSTRING `lb:16,ub:16`
	Resumecause    Resumecause
	Spare          utils.BITSTRING `lb:1,ub:1`
}
