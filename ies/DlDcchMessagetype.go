package ies

// DL-DCCH-MessageType ::= CHOICE
const (
	DlDcchMessagetypeChoiceNothing = iota
	DlDcchMessagetypeChoiceC1
	DlDcchMessagetypeChoiceMessageclassextension
)

type DlDcchMessagetype struct {
	Choice                uint64
	C1                    *DlDcchMessagetypeC1
	Messageclassextension *DlDcchMessagetypeMessageclassextension
}
