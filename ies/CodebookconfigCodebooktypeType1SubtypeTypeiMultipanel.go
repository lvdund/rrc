package ies

import "rrc/utils"

// CodebookConfig-codebookType-type1-subType-typeI-MultiPanel ::= SEQUENCE
type CodebookconfigCodebooktypeType1SubtypeTypeiMultipanel struct {
	NgN1N2        CodebookconfigCodebooktypeType1SubtypeTypeiMultipanelNgN1N2
	RiRestriction utils.BITSTRING `lb:4,ub:4`
}
