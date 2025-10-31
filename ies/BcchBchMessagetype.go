package ies

// BCCH-BCH-MessageType ::= CHOICE
const (
	BcchBchMessagetypeChoiceNothing = iota
	BcchBchMessagetypeChoiceMib
	BcchBchMessagetypeChoiceMessageclassextension
)

type BcchBchMessagetype struct {
	Choice                uint64
	Mib                   *Mib
	Messageclassextension *BcchBchMessagetypeMessageclassextension
}
