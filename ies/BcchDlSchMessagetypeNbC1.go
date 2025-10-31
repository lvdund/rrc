package ies

// BCCH-DL-SCH-MessageType-NB-c1 ::= CHOICE
const (
	BcchDlSchMessagetypeNbC1ChoiceNothing = iota
	BcchDlSchMessagetypeNbC1ChoiceSysteminformationR13
	BcchDlSchMessagetypeNbC1ChoiceSysteminformationblocktype1R13
)

type BcchDlSchMessagetypeNbC1 struct {
	Choice                         uint64
	SysteminformationR13           *SysteminformationNb
	Systeminformationblocktype1R13 *Systeminformationblocktype1Nb
}
