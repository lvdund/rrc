package ies

import "rrc/utils"

// MBMSInterestIndication-r11-IEs ::= SEQUENCE
type MbmsinterestindicationR11 struct {
	MbmsFreqlistR11          *CarrierfreqlistmbmsR11
	MbmsPriorityR11          *bool
	Latenoncriticalextension *utils.OCTETSTRING
	Noncriticalextension     *MbmsinterestindicationV1310
}
