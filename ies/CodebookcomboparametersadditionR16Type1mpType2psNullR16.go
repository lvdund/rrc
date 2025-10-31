package ies

import "rrc/utils"

// CodebookComboParametersAddition-r16-type1MP-Type2PS-null-r16 ::= SEQUENCE
type CodebookcomboparametersadditionR16Type1mpType2psNullR16 struct {
	SupportedcsiRsResourcelistaddR16 []utils.INTEGER `lb:1,ub:maxNrofCSIRsResourcesextR16`
}
