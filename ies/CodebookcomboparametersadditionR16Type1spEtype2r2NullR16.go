package ies

import "rrc/utils"

// CodebookComboParametersAddition-r16-type1SP-eType2R2-null-r16 ::= SEQUENCE
type CodebookcomboparametersadditionR16Type1spEtype2r2NullR16 struct {
	SupportedcsiRsResourcelistaddR16 []utils.INTEGER `lb:1,ub:maxNrofCSIRsResourcesextR16`
}
