package ies

import "rrc/utils"

// UplinkPowerControlDedicatedSCell-v1310 ::= SEQUENCE
type UplinkpowercontroldedicatedscellV1310 struct {
	P0UePucch                  utils.INTEGER `lb:0,ub:7`
	DeltatxdOffsetlistpucchR10 *DeltatxdOffsetlistpucchR10
}
