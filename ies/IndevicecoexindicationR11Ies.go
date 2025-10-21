package ies

import "rrc/utils"

// InDeviceCoexIndication-r11-IEs ::= SEQUENCE
type IndevicecoexindicationR11Ies struct {
	AffectedcarrierfreqlistR11 *AffectedcarrierfreqlistR11
	TdmAssistanceinfoR11       *TdmAssistanceinfoR11
	Latenoncriticalextension   *utils.OCTETSTRING
	Noncriticalextension       *IndevicecoexindicationV11d0Ies
}
