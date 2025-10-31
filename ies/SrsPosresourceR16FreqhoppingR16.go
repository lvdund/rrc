package ies

import "rrc/utils"

// SRS-PosResource-r16-freqHopping-r16 ::= SEQUENCE
// Extensible
type SrsPosresourceR16FreqhoppingR16 struct {
	CSrsR16 utils.INTEGER `lb:0,ub:63`
}
