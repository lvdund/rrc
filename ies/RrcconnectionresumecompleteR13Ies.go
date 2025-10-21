package ies

import "rrc/utils"

// RRCConnectionResumeComplete-r13-IEs ::= SEQUENCE
type RrcconnectionresumecompleteR13Ies struct {
	SelectedplmnIdentityR13     *utils.INTEGER
	DedicatedinfonasR13         *Dedicatedinfonas
	RlfInfoavailableR13         *utils.ENUMERATED
	LogmeasavailableR13         *utils.ENUMERATED
	ConnestfailinfoavailableR13 *utils.ENUMERATED
	MobilitystateR13            *utils.ENUMERATED
	MobilityhistoryavailR13     *utils.ENUMERATED
	LogmeasavailablembsfnR13    *utils.ENUMERATED
	Latenoncriticalextension    *utils.OCTETSTRING
	Noncriticalextension        *RrcconnectionresumecompleteV1530Ies
}
