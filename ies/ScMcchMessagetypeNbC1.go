package ies

// SC-MCCH-MessageType-NB-c1 ::= CHOICE
const (
	ScMcchMessagetypeNbC1ChoiceNothing = iota
	ScMcchMessagetypeNbC1ChoiceScptmconfigurationR14
)

type ScMcchMessagetypeNbC1 struct {
	Choice                uint64
	ScptmconfigurationR14 *ScptmconfigurationNbR14
}
