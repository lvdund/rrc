package ies

import "rrc/utils"

// NRDC-Parameters ::= SEQUENCE
type NrdcParameters struct {
	Measandmobparametersnrdc *Measandmobparametersmrdc
	Generalparametersnrdc    *GeneralparametersmrdcXddDiff
	FddAddUeNrdcCapabilities *UeMrdcCapabilityaddxddMode
	TddAddUeNrdcCapabilities *UeMrdcCapabilityaddxddMode
	Fr1AddUeNrdcCapabilities *UeMrdcCapabilityaddfrxMode
	Fr2AddUeNrdcCapabilities *UeMrdcCapabilityaddfrxMode
	Dummy2                   *utils.OCTETSTRING
	Dummy                    *NrdcParametersDummy
}
