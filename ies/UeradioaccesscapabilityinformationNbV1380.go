package ies

import "rrc/utils"

// UERadioAccessCapabilityInformation-NB-v1380-IEs ::= SEQUENCE
type UeradioaccesscapabilityinformationNbV1380 struct {
	Latenoncriticalextension *utils.OCTETSTRING
	Noncriticalextension     *UeradioaccesscapabilityinformationNbR14
}
