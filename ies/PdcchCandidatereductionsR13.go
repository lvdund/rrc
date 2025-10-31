package ies

// PDCCH-CandidateReductions-r13 ::= CHOICE
const (
	PdcchCandidatereductionsR13ChoiceNothing = iota
	PdcchCandidatereductionsR13ChoiceRelease
	PdcchCandidatereductionsR13ChoiceSetup
)

type PdcchCandidatereductionsR13 struct {
	Choice  uint64
	Release *struct{}
	Setup   *PdcchCandidatereductionsR13Setup
}
