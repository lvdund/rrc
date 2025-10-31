package ies

import "rrc/utils"

// CodebookParametersAddition-r16-etype2-r16-etype2R2-r16 ::= SEQUENCE
type CodebookparametersadditionR16Etype2R16Etype2r2R16 struct {
	SupportedcsiRsResourcelistaddR16 []utils.INTEGER `lb:1,ub:maxNrofCSIRsResourcesextR16`
}
