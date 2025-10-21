package ies

import "rrc/utils"

// UE-EUTRA-Capability-v10c0-IEs ::= SEQUENCE
type UeEutraCapabilityV10c0Ies struct {
	OtdoaPositioningcapabilitiesR10 *OtdoaPositioningcapabilitiesR10
	Noncriticalextension            *UeEutraCapabilityV10f0Ies
}
