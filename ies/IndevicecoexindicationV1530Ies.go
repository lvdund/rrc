package ies

import "rrc/utils"

// InDeviceCoexIndication-v1530-IEs ::= SEQUENCE
type IndevicecoexindicationV1530Ies struct {
	MrdcAssistanceinfoR15 *MrdcAssistanceinfoR15
	Noncriticalextension  *IndevicecoexindicationV1610Ies
}
