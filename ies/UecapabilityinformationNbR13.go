package ies

import "rrc/utils"

// UECapabilityInformation-NB-r13-IEs ::= SEQUENCE
type UecapabilityinformationNbR13 struct {
	UeCapabilityR13          UeCapabilityNbR13
	UeRadiopaginginfoR13     UeRadiopaginginfoNbR13
	Latenoncriticalextension *utils.OCTETSTRING
	Noncriticalextension     *UecapabilityinformationNbExtR14
}
