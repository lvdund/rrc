package ies

import "rrc/utils"

// UplinkPowerControlDedicated-v1130 ::= SEQUENCE
type UplinkpowercontroldedicatedV1130 struct {
	PsrsOffsetV1130              *utils.INTEGER `lb:0,ub:31`
	PsrsOffsetapV1130            *utils.INTEGER `lb:0,ub:31`
	DeltatxdOffsetlistpucchV1130 *DeltatxdOffsetlistpucchV1130
}
