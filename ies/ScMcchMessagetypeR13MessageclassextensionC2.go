package ies

// SC-MCCH-MessageType-r13-messageClassExtension-c2 ::= CHOICE
const (
	ScMcchMessagetypeR13MessageclassextensionC2ChoiceNothing = iota
	ScMcchMessagetypeR13MessageclassextensionC2ChoiceScptmconfigurationBrR14
	ScMcchMessagetypeR13MessageclassextensionC2ChoiceSpare
)

type ScMcchMessagetypeR13MessageclassextensionC2 struct {
	Choice                  uint64
	ScptmconfigurationBrR14 *ScptmconfigurationBrR14
	Spare                   *struct{}
}
