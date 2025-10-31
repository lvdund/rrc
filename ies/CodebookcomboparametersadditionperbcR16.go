package ies

import "rrc/utils"

// CodebookComboParametersAdditionPerBC-r16 ::= SEQUENCE
type CodebookcomboparametersadditionperbcR16 struct {
	Type1spType2NullR16      *[]utils.INTEGER `lb:1,ub:maxNrofCSIRsResourcesextR16`
	Type1spType2psNullR16    *[]utils.INTEGER `lb:1,ub:maxNrofCSIRsResourcesextR16`
	Type1spEtype2r1NullR16   *[]utils.INTEGER `lb:1,ub:maxNrofCSIRsResourcesextR16`
	Type1spEtype2r2NullR16   *[]utils.INTEGER `lb:1,ub:maxNrofCSIRsResourcesextR16`
	Type1spEtype2r1psNullR16 *[]utils.INTEGER `lb:1,ub:maxNrofCSIRsResourcesextR16`
	Type1spEtype2r2psNullR16 *[]utils.INTEGER `lb:1,ub:maxNrofCSIRsResourcesextR16`
	Type1spType2Type2psR16   *[]utils.INTEGER `lb:1,ub:maxNrofCSIRsResourcesextR16`
	Type1mpType2NullR16      *[]utils.INTEGER `lb:1,ub:maxNrofCSIRsResourcesextR16`
	Type1mpType2psNullR16    *[]utils.INTEGER `lb:1,ub:maxNrofCSIRsResourcesextR16`
	Type1mpEtype2r1NullR16   *[]utils.INTEGER `lb:1,ub:maxNrofCSIRsResourcesextR16`
	Type1mpEtype2r2NullR16   *[]utils.INTEGER `lb:1,ub:maxNrofCSIRsResourcesextR16`
	Type1mpEtype2r1psNullR16 *[]utils.INTEGER `lb:1,ub:maxNrofCSIRsResourcesextR16`
	Type1mpEtype2r2psNullR16 *[]utils.INTEGER `lb:1,ub:maxNrofCSIRsResourcesextR16`
	Type1mpType2Type2psR16   *[]utils.INTEGER `lb:1,ub:maxNrofCSIRsResourcesextR16`
}
