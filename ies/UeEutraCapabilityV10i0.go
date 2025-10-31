package ies

import "rrc/utils"

// UE-EUTRA-Capability-v10i0-IEs ::= SEQUENCE
type UeEutraCapabilityV10i0 struct {
	RfParametersV10i0        *RfParametersV10i0
	Latenoncriticalextension *utils.OCTETSTRING
	Noncriticalextension     *UeEutraCapabilityV11d0
}
