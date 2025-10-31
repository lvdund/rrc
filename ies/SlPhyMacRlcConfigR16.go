package ies

import "rrc/utils"

// SL-PHY-MAC-RLC-Config-r16 ::= SEQUENCE
type SlPhyMacRlcConfigR16 struct {
	SlScheduledconfigR16        *utils.Setuprelease[SlScheduledconfigR16]
	SlUeSelectedconfigR16       *utils.Setuprelease[SlUeSelectedconfigR16]
	SlFreqinfotoreleaselistR16  *[]SlFreqIdR16               `lb:1,ub:maxNrofFreqSLR16`
	SlFreqinfotoaddmodlistR16   *[]SlFreqconfigR16           `lb:1,ub:maxNrofFreqSLR16`
	SlRlcBearertoreleaselistR16 *[]SlRlcBearerconfigindexR16 `lb:1,ub:maxSLLcidR16`
	SlRlcBearertoaddmodlistR16  *[]SlRlcBearerconfigR16      `lb:1,ub:maxSLLcidR16`
	SlMaxnumconsecutivedtxR16   *SlPhyMacRlcConfigR16SlMaxnumconsecutivedtxR16
	SlCsiAcquisitionR16         *SlPhyMacRlcConfigR16SlCsiAcquisitionR16
	SlCsiSchedulingrequestidR16 *utils.Setuprelease[Schedulingrequestid]
	SlSsbPrioritynrR16          *utils.INTEGER `lb:0,ub:8`
	NetworkcontrolledsynctxR16  *SlPhyMacRlcConfigR16NetworkcontrolledsynctxR16
}
