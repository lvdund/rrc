package ies

import "rrc/utils"

// CandidateCellInfoList-r10 ::= SEQUENCE OF CandidateCellInfo-r10
// SIZE (1..maxFreq)
type CandidatecellinfolistR10 struct {
	Value []CandidatecellinfoR10
}
