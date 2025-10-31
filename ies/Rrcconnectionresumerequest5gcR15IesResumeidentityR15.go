package ies

// RRCConnectionResumeRequest-5GC-r15-IEs-resumeIdentity-r15 ::= CHOICE
const (
	Rrcconnectionresumerequest5gcR15IesResumeidentityR15ChoiceNothing = iota
	Rrcconnectionresumerequest5gcR15IesResumeidentityR15ChoiceFulliRntiR15
	Rrcconnectionresumerequest5gcR15IesResumeidentityR15ChoiceShortiRntiR15
)

type Rrcconnectionresumerequest5gcR15IesResumeidentityR15 struct {
	Choice        uint64
	FulliRntiR15  *IRntiR15
	ShortiRntiR15 *ShortiRntiR15
}
