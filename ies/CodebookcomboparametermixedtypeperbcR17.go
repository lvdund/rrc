package ies

import "rrc/utils"

// CodebookComboParameterMixedTypePerBC-r17 ::= SEQUENCE
type CodebookcomboparametermixedtypeperbcR17 struct {
	Type1spFetype2psNullR17         *[]utils.INTEGER `lb:1,ub:maxNrofCSIRsResourcesextR16`
	Type1spFetype2psM2r1NullR17     *[]utils.INTEGER `lb:1,ub:maxNrofCSIRsResourcesextR16`
	Type1spFetype2psM2r2NullR17     *[]utils.INTEGER `lb:1,ub:maxNrofCSIRsResourcesextR16`
	Type1spType2Fetype2PsM1R17      *[]utils.INTEGER `lb:1,ub:maxNrofCSIRsResourcesextR16`
	Type1spType2Fetype2PsM2r1R17    *[]utils.INTEGER `lb:1,ub:maxNrofCSIRsResourcesextR16`
	Type1spEtype2r1Fetype2PsM1R17   *[]utils.INTEGER `lb:1,ub:maxNrofCSIRsResourcesextR16`
	Type1spEtype2r1Fetype2PsM2r1R17 *[]utils.INTEGER `lb:1,ub:maxNrofCSIRsResourcesextR16`
	Type1mpFetype2psNullR17         *[]utils.INTEGER `lb:1,ub:maxNrofCSIRsResourcesextR16`
	Type1mpFetype2psM2r1NullR17     *[]utils.INTEGER `lb:1,ub:maxNrofCSIRsResourcesextR16`
	Type1mpFetype2psM2r2NullR17     *[]utils.INTEGER `lb:1,ub:maxNrofCSIRsResourcesextR16`
	Type1mpType2Fetype2PsM1R17      *[]utils.INTEGER `lb:1,ub:maxNrofCSIRsResourcesextR16`
	Type1mpType2Fetype2PsM2r1R17    *[]utils.INTEGER `lb:1,ub:maxNrofCSIRsResourcesextR16`
	Type1mpEtype2r1Fetype2PsM1R17   *[]utils.INTEGER `lb:1,ub:maxNrofCSIRsResourcesextR16`
	Type1mpEtype2r1Fetype2PsM2r1R17 *[]utils.INTEGER `lb:1,ub:maxNrofCSIRsResourcesextR16`
}
