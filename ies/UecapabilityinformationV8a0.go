package ies

import "rrc/utils"

// UECapabilityInformation-v8a0-IEs ::= SEQUENCE
type UecapabilityinformationV8a0 struct {
	Latenoncriticalextension *utils.OCTETSTRING
	Noncriticalextension     *UecapabilityinformationV1250
}
