package ies

import "rrc/utils"

// CodebookConfig-codebookType-type1-subType-typeI-SinglePanel-nrOfAntennaPorts-moreThanTwo ::= SEQUENCE
type CodebookconfigCodebooktypeType1SubtypeTypeiSinglepanelNrofantennaportsMorethantwo struct {
	N1N2                                        CodebookconfigCodebooktypeType1SubtypeTypeiSinglepanelNrofantennaportsMorethantwoN1N2
	TypeiSinglepanelCodebooksubsetrestrictionI2 *utils.BITSTRING `lb:16,ub:16`
}
