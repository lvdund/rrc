package ies

// UL-DCCH-MessageType ::= CHOICE
const (
	UlDcchMessagetypeChoiceNothing = iota
	UlDcchMessagetypeChoiceC1
	UlDcchMessagetypeChoiceMessageclassextension
)

type UlDcchMessagetype struct {
	Choice                uint64
	C1                    *UlDcchMessagetypeC1
	Messageclassextension *UlDcchMessagetypeMessageclassextension
}
