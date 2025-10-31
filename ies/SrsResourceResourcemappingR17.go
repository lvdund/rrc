package ies

import "rrc/utils"

// SRS-Resource-resourceMapping-r17 ::= SEQUENCE
type SrsResourceResourcemappingR17 struct {
	StartpositionR17    utils.INTEGER `lb:0,ub:13`
	NrofsymbolsR17      SrsResourceResourcemappingR17NrofsymbolsR17
	RepetitionfactorR17 SrsResourceResourcemappingR17RepetitionfactorR17
}
