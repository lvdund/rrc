package ies

// MCCH-MessageType-later ::= CHOICE
const (
	McchMessagetypeLaterChoiceNothing = iota
	McchMessagetypeLaterChoiceC2
	McchMessagetypeLaterChoiceMessageclassextension
)

type McchMessagetypeLater struct {
	Choice                uint64
	C2                    *McchMessagetypeLaterC2
	Messageclassextension *McchMessagetypeLaterMessageclassextension
}
