package ies

import "rrc/utils"

// UplinkConfig ::= SEQUENCE
// Extensible
type Uplinkconfig struct {
	Initialuplinkbwp                 *BwpUplinkdedicated
	UplinkbwpToreleaselist           *[]BwpId     `lb:1,ub:maxNrofBWPs`
	UplinkbwpToaddmodlist            *[]BwpUplink `lb:1,ub:maxNrofBWPs`
	FirstactiveuplinkbwpId           *BwpId
	PuschServingcellconfig           *utils.Setuprelease[PuschServingcellconfig]
	Carrierswitching                 *utils.Setuprelease[SrsCarrierswitching]
	Powerboostpi2bpsk                *utils.BOOLEAN                                `ext`
	UplinkchannelbwPerscsList        *[]ScsSpecificcarrier                         `lb:1,ub:maxSCSs,ext`
	EnableplRsUpdateforpuschSrsR16   *UplinkconfigEnableplRsUpdateforpuschSrsR16   `ext`
	EnabledefaultbeamplForpusch00R16 *UplinkconfigEnabledefaultbeamplForpusch00R16 `ext`
	EnabledefaultbeamplForpucchR16   *UplinkconfigEnabledefaultbeamplForpucchR16   `ext`
	EnabledefaultbeamplForsrsR16     *UplinkconfigEnabledefaultbeamplForsrsR16     `ext`
	UplinktxswitchingR16             *utils.Setuprelease[UplinktxswitchingR16]     `ext`
	MprPowerboostFr2R16              *utils.BOOLEAN                                `ext`
}
