package ies

// RLF-TimersAndConstantsSCG-r12 ::= CHOICE
// Extensible
const (
	RlfTimersandconstantsscgR12ChoiceNothing = iota
	RlfTimersandconstantsscgR12ChoiceRelease
	RlfTimersandconstantsscgR12ChoiceSetup
)

type RlfTimersandconstantsscgR12 struct {
	Choice  uint64
	Release *struct{}
	Setup   *RlfTimersandconstantsscgR12Setup
}
