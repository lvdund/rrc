package ies

import "rrc/utils"

// UE-EUTRA-Capability-v1560-IEs ::= SEQUENCE
type UeEutraCapabilityV1560 struct {
	PdcpParametersnrV1560            PdcpParametersnrV1560
	IratParametersnrV1560            IratParametersnrV1560
	AppliedcapabilityfiltercommonR15 *utils.OCTETSTRING
	FddAddUeEutraCapabilitiesV1560   UeEutraCapabilityaddxddModeV1560
	TddAddUeEutraCapabilitiesV1560   UeEutraCapabilityaddxddModeV1560
	Noncriticalextension             *UeEutraCapabilityV1570
}
