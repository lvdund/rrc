package ies

// BCCH-DL-SCH-MessageType-NB ::= CHOICE
const (
	BcchDlSchMessagetypeNbChoiceNothing = iota
	BcchDlSchMessagetypeNbChoiceC1
	BcchDlSchMessagetypeNbChoiceMessageclassextension
)

type BcchDlSchMessagetypeNb struct {
	Choice                uint64
	C1                    *BcchDlSchMessagetypeNbC1
	Messageclassextension *BcchDlSchMessagetypeNbMessageclassextension
}
