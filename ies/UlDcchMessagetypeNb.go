package ies

// UL-DCCH-MessageType-NB ::= CHOICE
const (
	UlDcchMessagetypeNbChoiceNothing = iota
	UlDcchMessagetypeNbChoiceC1
	UlDcchMessagetypeNbChoiceMessageclassextension
)

type UlDcchMessagetypeNb struct {
	Choice                uint64
	C1                    *UlDcchMessagetypeNbC1
	Messageclassextension *UlDcchMessagetypeNbMessageclassextension
}
