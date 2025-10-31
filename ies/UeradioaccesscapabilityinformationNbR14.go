package ies

import "rrc/utils"

// UERadioAccessCapabilityInformation-NB-r14-IEs ::= SEQUENCE
type UeradioaccesscapabilityinformationNbR14 struct {
	UeRadioaccesscapabilityinfoR14 *utils.OCTETSTRING
	Noncriticalextension           *UeradioaccesscapabilityinformationNbR14IesNoncriticalextension
}
