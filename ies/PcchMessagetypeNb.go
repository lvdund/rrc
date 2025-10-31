package ies

// PCCH-MessageType-NB ::= CHOICE
const (
	PcchMessagetypeNbChoiceNothing = iota
	PcchMessagetypeNbChoiceC1
	PcchMessagetypeNbChoiceMessageclassextension
)

type PcchMessagetypeNb struct {
	Choice                uint64
	C1                    *PcchMessagetypeNbC1
	Messageclassextension *PcchMessagetypeNbMessageclassextension
}
