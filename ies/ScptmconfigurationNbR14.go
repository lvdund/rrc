package ies

import "rrc/utils"

// SCPTMConfiguration-NB-r14 ::= SEQUENCE
type ScptmconfigurationNbR14 struct {
	ScMtchInfolistR14         ScMtchInfolistNbR14
	ScptmNeighbourcelllistR14 *ScptmNeighbourcelllistNbR14
	Latenoncriticalextension  *utils.OCTETSTRING
	Noncriticalextension      *ScptmconfigurationNbV1610
}
