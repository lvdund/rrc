package ies

import "rrc/utils"

// UE-NR-Capability-v15j0 ::= SEQUENCE
type UeNrCapabilityV15j0 struct {
	Latenoncriticalextension *utils.OCTETSTRING
	Noncriticalextension     *UeNrCapabilityV16a0
}
