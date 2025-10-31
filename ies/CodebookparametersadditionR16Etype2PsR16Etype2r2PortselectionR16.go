package ies

import "rrc/utils"

// CodebookParametersAddition-r16-etype2-PS-r16-etype2R2-PortSelection-r16 ::= SEQUENCE
type CodebookparametersadditionR16Etype2PsR16Etype2r2PortselectionR16 struct {
	SupportedcsiRsResourcelistaddR16 []utils.INTEGER `lb:1,ub:maxNrofCSIRsResourcesextR16`
}
