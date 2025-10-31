package ies

import "rrc/utils"

// CodebookConfig-codebookType-type1 ::= SEQUENCE
type CodebookconfigCodebooktypeType1 struct {
	Subtype      *CodebookconfigCodebooktypeType1Subtype
	Codebookmode utils.INTEGER `lb:0,ub:2`
}
