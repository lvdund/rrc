package ies

import "rrc/utils"

// CodebookConfig-codebookType-type2-subType-typeII ::= SEQUENCE
type CodebookconfigCodebooktypeType2SubtypeTypeii struct {
	N1N2Codebooksubsetrestriction CodebookconfigCodebooktypeType2SubtypeTypeiiN1N2Codebooksubsetrestriction
	TypeiiRiRestriction           utils.BITSTRING `lb:2,ub:2`
}
