package ies

import "rrc/utils"

// UERadioAccessCapabilityInformation-r8-IEs ::= SEQUENCE
type UeradioaccesscapabilityinformationR8Ies struct {
	UeRadioaccesscapabilityinfo utils.OCTETSTRING
	Noncriticalextension        *interface{}
}
