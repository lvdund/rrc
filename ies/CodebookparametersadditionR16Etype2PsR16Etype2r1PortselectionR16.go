package ies

import "rrc/utils"

// CodebookParametersAddition-r16-etype2-PS-r16-etype2R1-PortSelection-r16 ::= SEQUENCE
type CodebookparametersadditionR16Etype2PsR16Etype2r1PortselectionR16 struct {
	SupportedcsiRsResourcelistaddR16 []utils.INTEGER `lb:1,ub:maxNrofCSIRsResourcesextR16`
}
