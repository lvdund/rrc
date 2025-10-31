package ies

import "rrc/utils"

// CodebookParametersAdditionPerBC-r16 ::= SEQUENCE
type CodebookparametersadditionperbcR16 struct {
	Etype2r1R16              *[]utils.INTEGER `lb:1,ub:maxNrofCSIRsResourcesextR16`
	Etype2r2R16              *[]utils.INTEGER `lb:1,ub:maxNrofCSIRsResourcesextR16`
	Etype2r1PortselectionR16 *[]utils.INTEGER `lb:1,ub:maxNrofCSIRsResourcesextR16`
	Etype2r2PortselectionR16 *[]utils.INTEGER `lb:1,ub:maxNrofCSIRsResourcesextR16`
}
