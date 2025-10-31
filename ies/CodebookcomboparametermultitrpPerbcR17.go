package ies

import "rrc/utils"

// CodebookComboParameterMultiTRP-PerBC-r17 ::= SEQUENCE
type CodebookcomboparametermultitrpPerbcR17 struct {
	NcjtNullNull                    *[]utils.INTEGER `lb:1,ub:maxNrofCSIRsResourcesextR16`
	Ncjt1spNullNull                 *[]utils.INTEGER `lb:1,ub:maxNrofCSIRsResourcesextR16`
	NcjtType2NullR16                *[]utils.INTEGER `lb:1,ub:maxNrofCSIRsResourcesextR16`
	NcjtType2psNullR16              *[]utils.INTEGER `lb:1,ub:maxNrofCSIRsResourcesextR16`
	NcjtEtype2r1NullR16             *[]utils.INTEGER `lb:1,ub:maxNrofCSIRsResourcesextR16`
	NcjtEtype2r2NullR16             *[]utils.INTEGER `lb:1,ub:maxNrofCSIRsResourcesextR16`
	NcjtEtype2r1psNullR16           *[]utils.INTEGER `lb:1,ub:maxNrofCSIRsResourcesextR16`
	NcjtEtype2r2psNullR16           *[]utils.INTEGER `lb:1,ub:maxNrofCSIRsResourcesextR16`
	NcjtType2Type2psR16             *[]utils.INTEGER `lb:1,ub:maxNrofCSIRsResourcesextR16`
	Ncjt1spType2NullR16             *[]utils.INTEGER `lb:1,ub:maxNrofCSIRsResourcesextR16`
	Ncjt1spType2psNullR16           *[]utils.INTEGER `lb:1,ub:maxNrofCSIRsResourcesextR16`
	Ncjt1spEtype2r1NullR16          *[]utils.INTEGER `lb:1,ub:maxNrofCSIRsResourcesextR16`
	Ncjt1spEtype2r2NullR16          *[]utils.INTEGER `lb:1,ub:maxNrofCSIRsResourcesextR16`
	Ncjt1spEtype2r1psNullR16        *[]utils.INTEGER `lb:1,ub:maxNrofCSIRsResourcesextR16`
	Ncjt1spEtype2r2psNullR16        *[]utils.INTEGER `lb:1,ub:maxNrofCSIRsResourcesextR16`
	Ncjt1spType2Type2psR16          *[]utils.INTEGER `lb:1,ub:maxNrofCSIRsResourcesextR16`
	NcjtFetype2psNullR17            *[]utils.INTEGER `lb:1,ub:maxNrofCSIRsResourcesextR16`
	NcjtFetype2psM2r1NullR17        *[]utils.INTEGER `lb:1,ub:maxNrofCSIRsResourcesextR16`
	NcjtFetype2psM2r2NullR17        *[]utils.INTEGER `lb:1,ub:maxNrofCSIRsResourcesextR16`
	NcjtType2Fetype2PsM1R17         *[]utils.INTEGER `lb:1,ub:maxNrofCSIRsResourcesextR16`
	NcjtType2Fetype2PsM2r1R17       *[]utils.INTEGER `lb:1,ub:maxNrofCSIRsResourcesextR16`
	NcjtEtype2r1Fetype2PsM1R17      *[]utils.INTEGER `lb:1,ub:maxNrofCSIRsResourcesextR16`
	NcjtEtype2r1Fetype2PsM2r1R17    *[]utils.INTEGER `lb:1,ub:maxNrofCSIRsResourcesextR16`
	Ncjt1spFetype2psNullR17         *[]utils.INTEGER `lb:1,ub:maxNrofCSIRsResourcesextR16`
	Ncjt1spFetype2psM2r1NullR17     *[]utils.INTEGER `lb:1,ub:maxNrofCSIRsResourcesextR16`
	Ncjt1spFetype2psM2r2NullR1      *[]utils.INTEGER `lb:1,ub:maxNrofCSIRsResourcesextR16`
	Ncjt1spType2Fetype2PsM1R17      *[]utils.INTEGER `lb:1,ub:maxNrofCSIRsResourcesextR16`
	Ncjt1spType2Fetype2PsM2r1R17    *[]utils.INTEGER `lb:1,ub:maxNrofCSIRsResourcesextR16`
	Ncjt1spEtype2r1Fetype2PsM1R17   *[]utils.INTEGER `lb:1,ub:maxNrofCSIRsResourcesextR16`
	Ncjt1spEtype2r1Fetype2PsM2r1R17 *[]utils.INTEGER `lb:1,ub:maxNrofCSIRsResourcesextR16`
}
