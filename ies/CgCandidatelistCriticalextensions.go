package ies

// CG-CandidateList-criticalExtensions ::= CHOICE
const (
	CgCandidatelistCriticalextensionsChoiceNothing = iota
	CgCandidatelistCriticalextensionsChoiceC1
	CgCandidatelistCriticalextensionsChoiceCriticalextensionsfuture
)

type CgCandidatelistCriticalextensions struct {
	Choice                   uint64
	C1                       *CgCandidatelistCriticalextensionsC1
	Criticalextensionsfuture *CgCandidatelistCriticalextensionsCriticalextensionsfuture
}
