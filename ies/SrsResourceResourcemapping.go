package ies

import "rrc/utils"

// SRS-Resource-resourceMapping ::= SEQUENCE
type SrsResourceResourcemapping struct {
	Startposition    utils.INTEGER `lb:0,ub:5`
	Nrofsymbols      SrsResourceResourcemappingNrofsymbols
	Repetitionfactor SrsResourceResourcemappingRepetitionfactor
}
