package ies

import "rrc/utils"

// UE-EUTRA-Capability-v1660-IEs ::= SEQUENCE
type UeEutraCapabilityV1660Ies struct {
	IratParametersnrV1660 IratParametersnrV1660
	Noncriticalextension  *UeEutraCapabilityV1690Ies
}
