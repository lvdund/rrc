package ies

// PUCCH-ConfigDedicated-r13-pucch-NumRepetitionCE-r13-setup ::= CHOICE
const (
	PucchConfigdedicatedR13PucchNumrepetitionceR13SetupChoiceNothing = iota
	PucchConfigdedicatedR13PucchNumrepetitionceR13SetupChoiceModea
	PucchConfigdedicatedR13PucchNumrepetitionceR13SetupChoiceModeb
)

type PucchConfigdedicatedR13PucchNumrepetitionceR13Setup struct {
	Choice uint64
	Modea  *PucchConfigdedicatedR13PucchNumrepetitionceR13SetupModea
	Modeb  *PucchConfigdedicatedR13PucchNumrepetitionceR13SetupModeb
}
