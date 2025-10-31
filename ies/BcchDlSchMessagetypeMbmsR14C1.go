package ies

// BCCH-DL-SCH-MessageType-MBMS-r14-c1 ::= CHOICE
const (
	BcchDlSchMessagetypeMbmsR14C1ChoiceNothing = iota
	BcchDlSchMessagetypeMbmsR14C1ChoiceSysteminformationMbmsR14
	BcchDlSchMessagetypeMbmsR14C1ChoiceSysteminformationblocktype1MbmsR14
)

type BcchDlSchMessagetypeMbmsR14C1 struct {
	Choice                             uint64
	SysteminformationMbmsR14           *SysteminformationMbmsR14
	Systeminformationblocktype1MbmsR14 *Systeminformationblocktype1MbmsR14
}
