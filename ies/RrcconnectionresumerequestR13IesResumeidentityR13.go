package ies

import "rrc/utils"

// RRCConnectionResumeRequest-r13-IEs-resumeIdentity-r13 ::= CHOICE
const (
	RrcconnectionresumerequestR13IesResumeidentityR13ChoiceNothing = iota
	RrcconnectionresumerequestR13IesResumeidentityR13ChoiceResumeidR13
	RrcconnectionresumerequestR13IesResumeidentityR13ChoiceTruncatedresumeidR13
)

type RrcconnectionresumerequestR13IesResumeidentityR13 struct {
	Choice               uint64
	ResumeidR13          *ResumeidentityR13
	TruncatedresumeidR13 *utils.BITSTRING `lb:24,ub:24`
}
