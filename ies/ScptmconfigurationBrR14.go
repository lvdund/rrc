package ies

import "rrc/utils"

// SCPTMConfiguration-BR-r14 ::= SEQUENCE
type ScptmconfigurationBrR14 struct {
	ScMtchInfolistR14         ScMtchInfolistBrR14
	ScptmNeighbourcelllistR14 *ScptmNeighbourcelllistR13
	PBR14                     *utils.INTEGER `lb:0,ub:3`
	Latenoncriticalextension  *utils.OCTETSTRING
	Noncriticalextension      *ScptmconfigurationBrV1610
}
