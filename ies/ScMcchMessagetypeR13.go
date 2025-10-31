package ies

// SC-MCCH-MessageType-r13 ::= CHOICE
const (
	ScMcchMessagetypeR13ChoiceNothing = iota
	ScMcchMessagetypeR13ChoiceC1
	ScMcchMessagetypeR13ChoiceMessageclassextension
)

type ScMcchMessagetypeR13 struct {
	Choice                uint64
	C1                    *ScMcchMessagetypeR13C1
	Messageclassextension *ScMcchMessagetypeR13Messageclassextension
}
