package ies

// UL-CCCH-MessageType-NB ::= CHOICE
const (
	UlCcchMessagetypeNbChoiceNothing = iota
	UlCcchMessagetypeNbChoiceC1
	UlCcchMessagetypeNbChoiceMessageclassextension
)

type UlCcchMessagetypeNb struct {
	Choice                uint64
	C1                    *UlCcchMessagetypeNbC1
	Messageclassextension *UlCcchMessagetypeNbMessageclassextension
}
