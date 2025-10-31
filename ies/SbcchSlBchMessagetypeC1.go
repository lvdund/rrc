package ies

// SBCCH-SL-BCH-MessageType-c1 ::= CHOICE
const (
	SbcchSlBchMessagetypeC1ChoiceNothing = iota
	SbcchSlBchMessagetypeC1ChoiceMasterinformationblocksidelink
	SbcchSlBchMessagetypeC1ChoiceSpare1
)

type SbcchSlBchMessagetypeC1 struct {
	Choice                         uint64
	Masterinformationblocksidelink *Masterinformationblocksidelink
	Spare1                         *struct{}
}
