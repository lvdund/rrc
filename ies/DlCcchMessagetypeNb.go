package ies

// DL-CCCH-MessageType-NB ::= CHOICE
const (
	DlCcchMessagetypeNbChoiceNothing = iota
	DlCcchMessagetypeNbChoiceC1
	DlCcchMessagetypeNbChoiceMessageclassextension
)

type DlCcchMessagetypeNb struct {
	Choice                uint64
	C1                    *DlCcchMessagetypeNbC1
	Messageclassextension *DlCcchMessagetypeNbMessageclassextension
}
