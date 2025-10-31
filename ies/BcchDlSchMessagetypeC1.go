package ies

// BCCH-DL-SCH-MessageType-c1 ::= CHOICE
const (
	BcchDlSchMessagetypeC1ChoiceNothing = iota
	BcchDlSchMessagetypeC1ChoiceSysteminformation
	BcchDlSchMessagetypeC1ChoiceSysteminformationblocktype1
)

type BcchDlSchMessagetypeC1 struct {
	Choice                      uint64
	Systeminformation           *Systeminformation
	Systeminformationblocktype1 *Sib1
}
