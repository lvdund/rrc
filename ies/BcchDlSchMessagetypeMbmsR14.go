package ies

// BCCH-DL-SCH-MessageType-MBMS-r14 ::= CHOICE
const (
	BcchDlSchMessagetypeMbmsR14ChoiceNothing = iota
	BcchDlSchMessagetypeMbmsR14ChoiceC1
	BcchDlSchMessagetypeMbmsR14ChoiceMessageclassextension
)

type BcchDlSchMessagetypeMbmsR14 struct {
	Choice                uint64
	C1                    *BcchDlSchMessagetypeMbmsR14C1
	Messageclassextension *BcchDlSchMessagetypeMbmsR14Messageclassextension
}
