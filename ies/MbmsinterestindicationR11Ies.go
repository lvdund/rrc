package ies

import "rrc/utils"

// MBMSInterestIndication-r11-IEs ::= SEQUENCE
type MbmsinterestindicationR11Ies struct {
	MbmsFreqlistR11          *CarrierfreqlistmbmsR11
	MbmsPriorityR11          *utils.ENUMERATED
	Latenoncriticalextension *utils.OCTETSTRING
	Noncriticalextension     *MbmsinterestindicationV1310Ies
}
