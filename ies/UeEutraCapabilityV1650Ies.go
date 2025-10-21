package ies

import "rrc/utils"

// UE-EUTRA-Capability-v1650-IEs ::= SEQUENCE
type UeEutraCapabilityV1650Ies struct {
	OtherparametersV1650 *OtherParametersV1650
	Noncriticalextension *UeEutraCapabilityV1660Ies
}
