package ies

import "rrc/utils"

// SCPTMConfiguration-v1340 ::= SEQUENCE
type ScptmconfigurationV1340 struct {
	PBR13                *utils.INTEGER `lb:0,ub:3`
	Noncriticalextension *ScptmconfigurationV1340Noncriticalextension
}
