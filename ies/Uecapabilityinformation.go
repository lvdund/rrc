package ies

import "rrc/utils"

// UECapabilityInformation-IEs ::= SEQUENCE
type Uecapabilityinformation struct {
	UeCapabilityratContainerlist *UeCapabilityratContainerlist
	Latenoncriticalextension     *utils.OCTETSTRING
	Noncriticalextension         *UecapabilityinformationIesNoncriticalextension
}
