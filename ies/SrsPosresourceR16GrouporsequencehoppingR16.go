package ies

import "rrc/utils"

// SRS-PosResource-r16-groupOrSequenceHopping-r16 ::= ENUMERATED
type SrsPosresourceR16GrouporsequencehoppingR16 struct {
	Value utils.ENUMERATED
}

const (
	SrsPosresourceR16GrouporsequencehoppingR16EnumeratedNothing = iota
	SrsPosresourceR16GrouporsequencehoppingR16EnumeratedNeither
	SrsPosresourceR16GrouporsequencehoppingR16EnumeratedGrouphopping
	SrsPosresourceR16GrouporsequencehoppingR16EnumeratedSequencehopping
)
