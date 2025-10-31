package ies

import "rrc/utils"

// CG-Config-v1730-IEs ::= SEQUENCE
type CgConfigV1730 struct {
	Fr1CarriersScgR17    *utils.INTEGER `lb:0,ub:32`
	Fr2CarriersScgR17    *utils.INTEGER `lb:0,ub:32`
	Noncriticalextension *CgConfigV1730IesNoncriticalextension
}
