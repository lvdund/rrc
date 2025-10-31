package ies

// PCCH-MessageType ::= CHOICE
const (
	PcchMessagetypeChoiceNothing = iota
	PcchMessagetypeChoiceC1
	PcchMessagetypeChoiceMessageclassextension
)

type PcchMessagetype struct {
	Choice                uint64
	C1                    *PcchMessagetypeC1
	Messageclassextension *PcchMessagetypeMessageclassextension
}
