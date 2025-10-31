package ies

import "rrc/utils"

// SL-ConfigCommonNR-r16 ::= SEQUENCE
type SlConfigcommonnrR16 struct {
	SlFreqinfolistR16               *[]SlFreqconfigcommonR16 `lb:1,ub:maxNrofFreqSLR16`
	SlUeSelectedconfigR16           *SlUeSelectedconfigR16
	SlNrAnchorcarrierfreqlistR16    *SlNrAnchorcarrierfreqlistR16
	SlEutraAnchorcarrierfreqlistR16 *SlEutraAnchorcarrierfreqlistR16
	SlRadiobearerconfiglistR16      *[]SlRadiobearerconfigR16 `lb:1,ub:maxNrofSLRBR16`
	SlRlcBearerconfiglistR16        *[]SlRlcBearerconfigR16   `lb:1,ub:maxSLLcidR16`
	SlMeasconfigcommonR16           *SlMeasconfigcommonR16
	SlCsiAcquisitionR16             *SlConfigcommonnrR16SlCsiAcquisitionR16
	SlOffsetdfnR16                  *utils.INTEGER `lb:0,ub:1000`
	T400R16                         *SlConfigcommonnrR16T400R16
	SlMaxnumconsecutivedtxR16       *SlConfigcommonnrR16SlMaxnumconsecutivedtxR16
	SlSsbPrioritynrR16              *utils.INTEGER `lb:0,ub:8`
}
