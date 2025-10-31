package ies

// DL-CCCH-MessageType-messageClassExtension-c2 ::= CHOICE
const (
	DlCcchMessagetypeMessageclassextensionC2ChoiceNothing = iota
	DlCcchMessagetypeMessageclassextensionC2ChoiceRrcearlydatacompleteR15
	DlCcchMessagetypeMessageclassextensionC2ChoiceSpare3
	DlCcchMessagetypeMessageclassextensionC2ChoiceSpare2
	DlCcchMessagetypeMessageclassextensionC2ChoiceSpare1
)

type DlCcchMessagetypeMessageclassextensionC2 struct {
	Choice                  uint64
	RrcearlydatacompleteR15 *RrcearlydatacompleteR15
	Spare3                  *struct{}
	Spare2                  *struct{}
	Spare1                  *struct{}
}
