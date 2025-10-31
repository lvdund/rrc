package ies

// SC-MCCH-MessageType-r13-c1 ::= CHOICE
const (
	ScMcchMessagetypeR13C1ChoiceNothing = iota
	ScMcchMessagetypeR13C1ChoiceScptmconfigurationR13
)

type ScMcchMessagetypeR13C1 struct {
	Choice                uint64
	ScptmconfigurationR13 *ScptmconfigurationR13
}
