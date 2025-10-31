package ies

// MCCH-MessageType-r17 ::= CHOICE
const (
	McchMessagetypeR17ChoiceNothing = iota
	McchMessagetypeR17ChoiceC1
	McchMessagetypeR17ChoiceMessageclassextension
)

type McchMessagetypeR17 struct {
	Choice                uint64
	C1                    *McchMessagetypeR17C1
	Messageclassextension *McchMessagetypeR17Messageclassextension
}
