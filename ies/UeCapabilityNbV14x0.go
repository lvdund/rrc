package ies

import "rrc/utils"

// UE-Capability-NB-v14x0-IEs ::= SEQUENCE
type UeCapabilityNbV14x0 struct {
	Latenoncriticalextension *utils.OCTETSTRING
	Noncriticalextension     *UeCapabilityNbV1530
}
