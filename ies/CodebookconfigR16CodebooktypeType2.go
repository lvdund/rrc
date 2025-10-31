package ies

import "rrc/utils"

// CodebookConfig-r16-codebookType-type2 ::= SEQUENCE
type CodebookconfigR16CodebooktypeType2 struct {
	Subtype                             CodebookconfigR16CodebooktypeType2Subtype
	NumberofpmiSubbandspercqiSubbandR16 utils.INTEGER `lb:0,ub:2`
	ParamcombinationR16                 utils.INTEGER `lb:0,ub:8`
}
