package ies

import "rrc/utils"

// UERadioAccessCapabilityInformation-NB-v1380-IEs ::= SEQUENCE
type UeradioaccesscapabilityinformationNbV1380Ies struct {
	Latenoncriticalextension *utils.OCTETSTRING
	Noncriticalextension     *UeradioaccesscapabilityinformationNbR14Ies
}
