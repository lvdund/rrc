package ies

import "rrc/utils"

// UECapabilityInformation-r8-IEs ::= SEQUENCE
type UecapabilityinformationR8Ies struct {
	UeCapabilityratContainerlist UeCapabilityratContainerlist
	Noncriticalextension         *UecapabilityinformationV8a0Ies
}
