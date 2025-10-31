package ies

import "rrc/utils"

// UERadioAccessCapabilityInformation-NB-IEs ::= SEQUENCE
type UeradioaccesscapabilityinformationNb struct {
	UeRadioaccesscapabilityinfoR13 utils.OCTETSTRING
	Noncriticalextension           *UeradioaccesscapabilityinformationNbV1380
}
