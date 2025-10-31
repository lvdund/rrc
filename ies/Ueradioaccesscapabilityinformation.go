package ies

import "rrc/utils"

// UERadioAccessCapabilityInformation-IEs ::= SEQUENCE
type Ueradioaccesscapabilityinformation struct {
	UeRadioaccesscapabilityinfo utils.OCTETSTRING
	Noncriticalextension        *UeradioaccesscapabilityinformationIesNoncriticalextension
}
