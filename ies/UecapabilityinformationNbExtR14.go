package ies

import "rrc/utils"

// UECapabilityInformation-NB-Ext-r14-IEs ::= SEQUENCE
type UecapabilityinformationNbExtR14 struct {
	UeCapabilityContainerextR14 utils.OCTETSTRING
	Noncriticalextension        *UecapabilityinformationNbExtR14IesNoncriticalextension
}
