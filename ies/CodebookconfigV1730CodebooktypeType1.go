package ies

import "rrc/utils"

// CodebookConfig-v1730-codebookType-type1 ::= SEQUENCE
type CodebookconfigV1730CodebooktypeType1 struct {
	Codebookmode *utils.INTEGER `lb:0,ub:2`
}
