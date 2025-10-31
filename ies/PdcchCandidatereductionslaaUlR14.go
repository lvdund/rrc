package ies

// PDCCH-CandidateReductionsLAA-UL-r14 ::= CHOICE
const (
	PdcchCandidatereductionslaaUlR14ChoiceNothing = iota
	PdcchCandidatereductionslaaUlR14ChoiceRelease
	PdcchCandidatereductionslaaUlR14ChoiceSetup
)

type PdcchCandidatereductionslaaUlR14 struct {
	Choice  uint64
	Release *struct{}
	Setup   *PdcchCandidatereductionslaaUlR14Setup
}
