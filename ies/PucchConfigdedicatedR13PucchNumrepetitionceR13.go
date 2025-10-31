package ies

// PUCCH-ConfigDedicated-r13-pucch-NumRepetitionCE-r13 ::= CHOICE
const (
	PucchConfigdedicatedR13PucchNumrepetitionceR13ChoiceNothing = iota
	PucchConfigdedicatedR13PucchNumrepetitionceR13ChoiceRelease
	PucchConfigdedicatedR13PucchNumrepetitionceR13ChoiceSetup
)

type PucchConfigdedicatedR13PucchNumrepetitionceR13 struct {
	Choice  uint64
	Release *struct{}
	Setup   *PucchConfigdedicatedR13PucchNumrepetitionceR13Setup
}
