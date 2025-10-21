package ies

import "rrc/utils"

// UE-EUTRA-Capability-v1570-IEs ::= SEQUENCE
type UeEutraCapabilityV1570Ies struct {
	RfParametersV1570     *RfParametersV1570
	IratParametersnrV1570 *IratParametersnrV1570
	Noncriticalextension  *UeEutraCapabilityV15a0Ies
}
