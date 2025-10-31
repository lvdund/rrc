package ies

import "rrc/utils"

// RRCResumeComplete-IEs ::= SEQUENCE
type Rrcresumecomplete struct {
	DedicatednasMessage       *DedicatednasMessage
	SelectedplmnIdentity      *utils.INTEGER `lb:0,ub:maxPLMN`
	Uplinktxdirectcurrentlist *Uplinktxdirectcurrentlist
	Latenoncriticalextension  *utils.OCTETSTRING
	Noncriticalextension      *RrcresumecompleteV1610
}
