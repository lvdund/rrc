package ies

import "rrc/utils"

// UE-EUTRA-Capability-v16c0-IEs ::= SEQUENCE
type UeEutraCapabilityV16c0Ies struct {
	MeasparametersV16c0      MeasparametersV16c0
	Latenoncriticalextension *utils.OCTETSTRING
	Noncriticalextension     *interface{}
}
