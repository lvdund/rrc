package ies

// PUCCH-ConfigDedicated-r13-ackNackRepetition-r13 ::= CHOICE
const (
	PucchConfigdedicatedR13AcknackrepetitionR13ChoiceNothing = iota
	PucchConfigdedicatedR13AcknackrepetitionR13ChoiceRelease
	PucchConfigdedicatedR13AcknackrepetitionR13ChoiceSetup
)

type PucchConfigdedicatedR13AcknackrepetitionR13 struct {
	Choice  uint64
	Release *struct{}
	Setup   *PucchConfigdedicatedR13AcknackrepetitionR13Setup
}
