package ies

import "rrc/utils"

// CodebookParametersfetype2PerBC-r17 ::= SEQUENCE
type Codebookparametersfetype2perbcR17 struct {
	Fetype2basicR17 []utils.INTEGER  `lb:1,ub:maxNrofCSIRsResourcesextR16`
	Fetype2r1R17    *[]utils.INTEGER `lb:1,ub:maxNrofCSIRsResourcesextR17`
	Fetype2r2R17    *[]utils.INTEGER `lb:1,ub:maxNrofCSIRsResourcesextR17`
}
