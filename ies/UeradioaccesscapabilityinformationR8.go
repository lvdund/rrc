package ies

import "rrc/utils"

// UERadioAccessCapabilityInformation-r8-IEs ::= SEQUENCE
type UeradioaccesscapabilityinformationR8 struct {
	UeRadioaccesscapabilityinfo utils.OCTETSTRING
	Noncriticalextension        *UeradioaccesscapabilityinformationR8IesNoncriticalextension
}
