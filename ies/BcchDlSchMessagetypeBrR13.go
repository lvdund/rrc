package ies

// BCCH-DL-SCH-MessageType-BR-r13 ::= CHOICE
const (
	BcchDlSchMessagetypeBrR13ChoiceNothing = iota
	BcchDlSchMessagetypeBrR13ChoiceC1
	BcchDlSchMessagetypeBrR13ChoiceMessageclassextension
)

type BcchDlSchMessagetypeBrR13 struct {
	Choice                uint64
	C1                    *BcchDlSchMessagetypeBrR13C1
	Messageclassextension *BcchDlSchMessagetypeBrR13Messageclassextension
}
