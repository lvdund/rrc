package ies

// CG-CandidateList-criticalExtensions-c1 ::= CHOICE
const (
	CgCandidatelistCriticalextensionsC1ChoiceNothing = iota
	CgCandidatelistCriticalextensionsC1ChoiceCgCandidatelistR17
	CgCandidatelistCriticalextensionsC1ChoiceSpare3
	CgCandidatelistCriticalextensionsC1ChoiceSpare2
	CgCandidatelistCriticalextensionsC1ChoiceSpare1
)

type CgCandidatelistCriticalextensionsC1 struct {
	Choice             uint64
	CgCandidatelistR17 *CgCandidatelistR17
	Spare3             *struct{}
	Spare2             *struct{}
	Spare1             *struct{}
}
