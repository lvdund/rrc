package ies

import "rrc/utils"

// UE-MRDC-Capability-v1560 ::= SEQUENCE
type UeMrdcCapabilityV1560 struct {
	Receivedfilters               *utils.OCTETSTRING
	MeasandmobparametersmrdcV1560 *MeasandmobparametersmrdcV1560
	FddAddUeMrdcCapabilitiesV1560 *UeMrdcCapabilityaddxddModeV1560
	TddAddUeMrdcCapabilitiesV1560 *UeMrdcCapabilityaddxddModeV1560
	Noncriticalextension          *UeMrdcCapabilityV1610
}
