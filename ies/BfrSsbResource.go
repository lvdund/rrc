package ies

import "rrc/utils"

// BFR-SSB-Resource ::= SEQUENCE
// Extensible
type BfrSsbResource struct {
	Ssb             SsbIndex
	RaPreambleindex utils.INTEGER `lb:0,ub:63`
}
