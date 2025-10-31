package ies

import "rrc/utils"

// RRCConnectionResumeComplete-r13-IEs ::= SEQUENCE
type RrcconnectionresumecompleteR13 struct {
	SelectedplmnIdentityR13     *utils.INTEGER `lb:0,ub:maxPLMNR11`
	DedicatedinfonasR13         *Dedicatedinfonas
	RlfInfoavailableR13         *bool
	LogmeasavailableR13         *bool
	ConnestfailinfoavailableR13 *bool
	MobilitystateR13            *RrcconnectionresumecompleteR13IesMobilitystateR13
	MobilityhistoryavailR13     *bool
	LogmeasavailablembsfnR13    *bool
	Latenoncriticalextension    *utils.OCTETSTRING
	Noncriticalextension        *RrcconnectionresumecompleteV1530
}
