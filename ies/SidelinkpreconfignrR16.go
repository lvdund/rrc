package ies

import "rrc/utils"

// SidelinkPreconfigNR-r16 ::= SEQUENCE
// Extensible
type SidelinkpreconfignrR16 struct {
	SlPreconfigfreqinfolistR16               *[]SlFreqconfigcommonR16 `lb:1,ub:maxNrofFreqSLR16`
	SlPreconfignrAnchorcarrierfreqlistR16    *SlNrAnchorcarrierfreqlistR16
	SlPreconfigeutraAnchorcarrierfreqlistR16 *SlEutraAnchorcarrierfreqlistR16
	SlRadiobearerpreconfiglistR16            *[]SlRadiobearerconfigR16 `lb:1,ub:maxNrofSLRBR16`
	SlRlcBearerpreconfiglistR16              *[]SlRlcBearerconfigR16   `lb:1,ub:maxSLLcidR16`
	SlMeaspreconfigR16                       *SlMeasconfigcommonR16
	SlOffsetdfnR16                           *utils.INTEGER `lb:0,ub:1000`
	T400R16                                  *SidelinkpreconfignrR16T400R16
	SlMaxnumconsecutivedtxR16                *SidelinkpreconfignrR16SlMaxnumconsecutivedtxR16
	SlSsbPrioritynrR16                       *utils.INTEGER `lb:0,ub:8`
	SlPreconfiggeneralR16                    *SlPreconfiggeneralR16
	SlUeSelectedpreconfigR16                 *SlUeSelectedconfigR16
	SlCsiAcquisitionR16                      *SidelinkpreconfignrR16SlCsiAcquisitionR16
	SlRohcProfilesR16                        *SlRohcProfilesR16
	SlMaxcidR16                              utils.INTEGER        `lb:0,ub:16383`
	SlDrxPreconfiggcBcR17                    *SlDrxConfiggcBcR17  `ext`
	SlTxprofilelistR17                       *SlTxprofilelistR17  `ext`
	SlPreconfigdiscconfigR17                 *SlRemoteueConfigR17 `ext`
}
