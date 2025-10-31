package ies

import "rrc/utils"

// CG-ConfigInfo-v1730-IEs ::= SEQUENCE
type CgConfiginfoV1730 struct {
	Fr1CarriersMcgR17    *utils.INTEGER `lb:0,ub:32`
	Fr2CarriersMcgR17    *utils.INTEGER `lb:0,ub:32`
	Noncriticalextension *CgConfiginfoV1730IesNoncriticalextension
}
