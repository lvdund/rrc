package ies

// SBCCH-SL-BCH-MessageType ::= CHOICE
const (
	SbcchSlBchMessagetypeChoiceNothing = iota
	SbcchSlBchMessagetypeChoiceC1
	SbcchSlBchMessagetypeChoiceMessageclassextension
)

type SbcchSlBchMessagetype struct {
	Choice                uint64
	C1                    *SbcchSlBchMessagetypeC1
	Messageclassextension *SbcchSlBchMessagetypeMessageclassextension
}
