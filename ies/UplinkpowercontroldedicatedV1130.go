package ies

import "rrc/utils"

// UplinkPowerControlDedicated-v1130 ::= SEQUENCE
type UplinkpowercontroldedicatedV1130 struct {
	PsrsOffsetV1130              *utils.INTEGER
	PsrsOffsetapV1130            *utils.INTEGER
	DeltatxdOffsetlistpucchV1130 *DeltatxdOffsetlistpucchV1130
}
