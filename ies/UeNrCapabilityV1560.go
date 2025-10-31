package ies

import "rrc/utils"

// UE-NR-Capability-v1560 ::= SEQUENCE
type UeNrCapabilityV1560 struct {
	NrdcParameters       *NrdcParameters
	Receivedfilters      *utils.OCTETSTRING
	Noncriticalextension *UeNrCapabilityV1570
}
