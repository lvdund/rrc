package ies

import "rrc/utils"

// CodebookComboParametersAddition-r16-type1SP-eType2R2PS-null-r16 ::= SEQUENCE
type CodebookcomboparametersadditionR16Type1spEtype2r2psNullR16 struct {
	SupportedcsiRsResourcelistaddR16 []utils.INTEGER `lb:1,ub:maxNrofCSIRsResourcesextR16`
}
