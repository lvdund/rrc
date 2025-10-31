package ies

// DL-CCCH-MessageType ::= CHOICE
const (
	DlCcchMessagetypeChoiceNothing = iota
	DlCcchMessagetypeChoiceC1
	DlCcchMessagetypeChoiceMessageclassextension
)

type DlCcchMessagetype struct {
	Choice                uint64
	C1                    *DlCcchMessagetypeC1
	Messageclassextension *DlCcchMessagetypeMessageclassextension
}
