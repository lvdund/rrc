package ies

import "rrc/utils"

// CodebookConfig-r16-codebookType-type2-subType-typeII-r16 ::= SEQUENCE
type CodebookconfigR16CodebooktypeType2SubtypeTypeiiR16 struct {
	N1N2CodebooksubsetrestrictionR16 CodebookconfigR16CodebooktypeType2SubtypeTypeiiR16N1N2CodebooksubsetrestrictionR16
	TypeiiRiRestrictionR16           utils.BITSTRING `lb:4,ub:4`
}
