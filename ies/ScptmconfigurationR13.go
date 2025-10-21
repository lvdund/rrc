package ies

import "rrc/utils"

// SCPTMConfiguration-r13 ::= SEQUENCE
type ScptmconfigurationR13 struct {
	ScMtchInfolistR13         ScMtchInfolistR13
	ScptmNeighbourcelllistR13 *ScptmNeighbourcelllistR13
	Latenoncriticalextension  *utils.OCTETSTRING
	Noncriticalextension      *ScptmconfigurationV1340
}
