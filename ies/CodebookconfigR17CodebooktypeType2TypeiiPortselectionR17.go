package ies

import "rrc/utils"

// CodebookConfig-r17-codebookType-type2-typeII-PortSelection-r17 ::= SEQUENCE
type CodebookconfigR17CodebooktypeType2TypeiiPortselectionR17 struct {
	ParamcombinationR17                 utils.INTEGER `lb:0,ub:8`
	ValueofnR17                         *CodebookconfigR17CodebooktypeType2TypeiiPortselectionR17ValueofnR17
	NumberofpmiSubbandspercqiSubbandR17 *utils.INTEGER  `lb:0,ub:2`
	TypeiiPortselectionriRestrictionR17 utils.BITSTRING `lb:4,ub:4`
}
