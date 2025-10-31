package ies

// UL-CCCH-MessageType ::= CHOICE
const (
	UlCcchMessagetypeChoiceNothing = iota
	UlCcchMessagetypeChoiceC1
	UlCcchMessagetypeChoiceMessageclassextension
)

type UlCcchMessagetype struct {
	Choice                uint64
	C1                    *UlCcchMessagetypeC1
	Messageclassextension *UlCcchMessagetypeMessageclassextension
}
