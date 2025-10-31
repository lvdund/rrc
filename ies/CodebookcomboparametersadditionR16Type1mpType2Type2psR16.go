package ies

import "rrc/utils"

// CodebookComboParametersAddition-r16-type1MP-Type2-Type2PS-r16 ::= SEQUENCE
type CodebookcomboparametersadditionR16Type1mpType2Type2psR16 struct {
	SupportedcsiRsResourcelistaddR16 []utils.INTEGER `lb:1,ub:maxNrofCSIRsResourcesextR16`
}
