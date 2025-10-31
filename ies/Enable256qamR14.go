package ies

// Enable256QAM-r14 ::= CHOICE
const (
	Enable256qamR14ChoiceNothing = iota
	Enable256qamR14ChoiceRelease
	Enable256qamR14ChoiceSetup
)

type Enable256qamR14 struct {
	Choice  uint64
	Release *struct{}
	Setup   *Enable256qamR14Setup
}
