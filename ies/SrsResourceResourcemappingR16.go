package ies

import "rrc/utils"

// SRS-Resource-resourceMapping-r16 ::= SEQUENCE
type SrsResourceResourcemappingR16 struct {
	StartpositionR16    utils.INTEGER `lb:0,ub:13`
	NrofsymbolsR16      SrsResourceResourcemappingR16NrofsymbolsR16
	RepetitionfactorR16 SrsResourceResourcemappingR16RepetitionfactorR16
}
