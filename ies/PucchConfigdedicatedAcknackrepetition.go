package ies

// PUCCH-ConfigDedicated-ackNackRepetition ::= CHOICE
const (
	PucchConfigdedicatedAcknackrepetitionChoiceNothing = iota
	PucchConfigdedicatedAcknackrepetitionChoiceRelease
	PucchConfigdedicatedAcknackrepetitionChoiceSetup
)

type PucchConfigdedicatedAcknackrepetition struct {
	Choice  uint64
	Release *struct{}
	Setup   *PucchConfigdedicatedAcknackrepetitionSetup
}
