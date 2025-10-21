package ies

import "rrc/utils"

// InDeviceCoexIndication-v1360-IEs ::= SEQUENCE
type IndevicecoexindicationV1360Ies struct {
	HardwaresharingproblemR13 *utils.ENUMERATED
	Noncriticalextension      *IndevicecoexindicationV1530Ies
}
