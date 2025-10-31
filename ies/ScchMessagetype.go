package ies

// SCCH-MessageType ::= CHOICE
const (
	ScchMessagetypeChoiceNothing = iota
	ScchMessagetypeChoiceC1
	ScchMessagetypeChoiceMessageclassextension
)

type ScchMessagetype struct {
	Choice                uint64
	C1                    *ScchMessagetypeC1
	Messageclassextension *ScchMessagetypeMessageclassextension
}
