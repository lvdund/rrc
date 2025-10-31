package ies

// RLF-TimersAndConstantsMCG-Failure-r16 ::= CHOICE
// Extensible
const (
	RlfTimersandconstantsmcgFailureR16ChoiceNothing = iota
	RlfTimersandconstantsmcgFailureR16ChoiceRelease
	RlfTimersandconstantsmcgFailureR16ChoiceSetup
)

type RlfTimersandconstantsmcgFailureR16 struct {
	Choice  uint64
	Release *struct{}
	Setup   *RlfTimersandconstantsmcgFailureR16Setup
}
