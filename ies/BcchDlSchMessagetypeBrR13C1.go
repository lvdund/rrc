package ies

// BCCH-DL-SCH-MessageType-BR-r13-c1 ::= CHOICE
const (
	BcchDlSchMessagetypeBrR13C1ChoiceNothing = iota
	BcchDlSchMessagetypeBrR13C1ChoiceSysteminformationBrR13
	BcchDlSchMessagetypeBrR13C1ChoiceSysteminformationblocktype1BrR13
)

type BcchDlSchMessagetypeBrR13C1 struct {
	Choice                           uint64
	SysteminformationBrR13           *SysteminformationBrR13
	Systeminformationblocktype1BrR13 *Systeminformationblocktype1BrR13
}
