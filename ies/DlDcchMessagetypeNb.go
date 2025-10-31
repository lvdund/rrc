package ies

// DL-DCCH-MessageType-NB ::= CHOICE
const (
	DlDcchMessagetypeNbChoiceNothing = iota
	DlDcchMessagetypeNbChoiceC1
	DlDcchMessagetypeNbChoiceMessageclassextension
)

type DlDcchMessagetypeNb struct {
	Choice                uint64
	C1                    *DlDcchMessagetypeNbC1
	Messageclassextension *DlDcchMessagetypeNbMessageclassextension
}
