package ies

// SC-MCCH-MessageType-NB ::= CHOICE
const (
	ScMcchMessagetypeNbChoiceNothing = iota
	ScMcchMessagetypeNbChoiceC1
	ScMcchMessagetypeNbChoiceMessageclassextension
)

type ScMcchMessagetypeNb struct {
	Choice                uint64
	C1                    *ScMcchMessagetypeNbC1
	Messageclassextension *ScMcchMessagetypeNbMessageclassextension
}
