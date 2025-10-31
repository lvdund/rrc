package ies

import "rrc/utils"

// UE-EUTRA-Capability-v9h0-IEs ::= SEQUENCE
type UeEutraCapabilityV9h0 struct {
	InterratParametersutraV9h0 *IratParametersutraV9h0
	Latenoncriticalextension   *utils.OCTETSTRING
	Noncriticalextension       *UeEutraCapabilityV10c0
}
