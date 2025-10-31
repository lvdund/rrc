package ies

import "rrc/utils"

// CodebookConfig-r17-codebookType-type1 ::= SEQUENCE
type CodebookconfigR17CodebooktypeType1 struct {
	TypeiSinglepanelGroup1R17            *CodebookconfigR17CodebooktypeType1TypeiSinglepanelGroup1R17
	TypeiSinglepanelGroup2R17            *CodebookconfigR17CodebooktypeType1TypeiSinglepanelGroup2R17
	TypeiSinglepanelRiRestrictionstrpR17 *utils.BITSTRING `lb:8,ub:8`
	TypeiSinglepanelRiRestrictionsdmR17  *utils.BITSTRING `lb:4,ub:4`
}
