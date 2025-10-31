package ies

import "rrc/utils"

// SRS-PosResource-r16-resourceMapping-r16 ::= SEQUENCE
type SrsPosresourceR16ResourcemappingR16 struct {
	StartpositionR16 utils.INTEGER `lb:0,ub:13`
	NrofsymbolsR16   SrsPosresourceR16ResourcemappingR16NrofsymbolsR16
}
