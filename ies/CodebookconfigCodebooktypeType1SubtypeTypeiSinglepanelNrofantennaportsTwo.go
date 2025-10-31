package ies

import "rrc/utils"

// CodebookConfig-codebookType-type1-subType-typeI-SinglePanel-nrOfAntennaPorts-two ::= SEQUENCE
type CodebookconfigCodebooktypeType1SubtypeTypeiSinglepanelNrofantennaportsTwo struct {
	TwotxCodebooksubsetrestriction utils.BITSTRING `lb:6,ub:6`
}
