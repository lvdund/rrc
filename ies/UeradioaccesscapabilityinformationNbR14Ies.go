package ies

import "rrc/utils"

// UERadioAccessCapabilityInformation-NB-r14-IEs ::= SEQUENCE
type UeradioaccesscapabilityinformationNbR14Ies struct {
	UeRadioaccesscapabilityinfoR14 *utils.OCTETSTRING
	Noncriticalextension           *interface{}
}
