package ies

import "rrc/utils"

// CodebookConfig-codebookType-type1-subType-typeI-SinglePanel ::= SEQUENCE
type CodebookconfigCodebooktypeType1SubtypeTypeiSinglepanel struct {
	Nrofantennaports              *CodebookconfigCodebooktypeType1SubtypeTypeiSinglepanelNrofantennaports
	TypeiSinglepanelRiRestriction utils.BITSTRING `lb:8,ub:8`
}
