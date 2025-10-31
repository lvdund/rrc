package ies

import "rrc/utils"

// MBSInterestIndication-r17-IEs ::= SEQUENCE
type MbsinterestindicationR17 struct {
	MbsFreqlistR17           *CarrierfreqlistmbsR17
	MbsPriorityR17           *utils.BOOLEAN
	MbsServicelistR17        *MbsServicelistR17
	Latenoncriticalextension *utils.OCTETSTRING
	Noncriticalextension     *MbsinterestindicationR17IesNoncriticalextension
}
