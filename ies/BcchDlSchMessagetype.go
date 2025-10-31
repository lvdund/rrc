package ies

// BCCH-DL-SCH-MessageType ::= CHOICE
const (
	BcchDlSchMessagetypeChoiceNothing = iota
	BcchDlSchMessagetypeChoiceC1
	BcchDlSchMessagetypeChoiceMessageclassextension
)

type BcchDlSchMessagetype struct {
	Choice                uint64
	C1                    *BcchDlSchMessagetypeC1
	Messageclassextension *BcchDlSchMessagetypeMessageclassextension
}
