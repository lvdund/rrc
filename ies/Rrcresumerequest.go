package ies

import "rrc/utils"

// RRCResumeRequest-IEs ::= SEQUENCE
type Rrcresumerequest struct {
	Resumeidentity ShortiRntiValue
	ResumemacI     utils.BITSTRING `lb:16,ub:16`
	Resumecause    Resumecause
	Spare          utils.BITSTRING `lb:1,ub:1`
}
