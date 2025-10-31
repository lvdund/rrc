package ies

import "rrc/utils"

// MBSBroadcastConfiguration-r17-IEs ::= SEQUENCE
type MbsbroadcastconfigurationR17 struct {
	MbsSessioninfolistR17       *MbsSessioninfolistR17
	MbsNeighbourcelllistR17     *MbsNeighbourcelllistR17
	DrxConfigptmListR17         *[]DrxConfigptmR17 `lb:1,ub:maxNrofDRXConfigptmR17`
	PdschConfigmtchR17          *PdschConfigbroadcastR17
	MtchSsbMappingwindowlistR17 *MtchSsbMappingwindowlistR17
	Latenoncriticalextension    *utils.OCTETSTRING
	Noncriticalextension        *MbsbroadcastconfigurationR17IesNoncriticalextension
}
