package ies

import "rrc/utils"

// UplinkPowerControlDedicated-v1020 ::= SEQUENCE
type UplinkpowercontroldedicatedV1020 struct {
	DeltatxdOffsetlistpucchR10 *DeltatxdOffsetlistpucchR10
	PsrsOffsetapR10            *utils.INTEGER `lb:0,ub:15`
}
