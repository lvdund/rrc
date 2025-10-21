package ies

import "rrc/utils"

// UECapabilityInformation-NB-Ext-r14-IEs ::= SEQUENCE
type UecapabilityinformationNbExtR14Ies struct {
	UeCapabilityContainerextR14 utils.OCTETSTRING
	Noncriticalextension        *interface{}
}
