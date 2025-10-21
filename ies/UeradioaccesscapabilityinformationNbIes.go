package ies

import "rrc/utils"

// UERadioAccessCapabilityInformation-NB-IEs ::= SEQUENCE
type UeradioaccesscapabilityinformationNbIes struct {
	UeRadioaccesscapabilityinfoR13 utils.OCTETSTRING
	Noncriticalextension           *UeradioaccesscapabilityinformationNbV1380Ies
}
