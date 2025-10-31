package ies

import "rrc/utils"

// CodebookComboParametersAddition-r16-type1MP-eType2R1PS-null-r16 ::= SEQUENCE
type CodebookcomboparametersadditionR16Type1mpEtype2r1psNullR16 struct {
	SupportedcsiRsResourcelistaddR16 []utils.INTEGER `lb:1,ub:maxNrofCSIRsResourcesextR16`
}
