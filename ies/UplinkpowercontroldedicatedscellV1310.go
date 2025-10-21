package ies

import "rrc/utils"

// UplinkPowerControlDedicatedSCell-v1310 ::= SEQUENCE
type UplinkpowercontroldedicatedscellV1310 struct {
	P0UePucch                  utils.INTEGER
	DeltatxdOffsetlistpucchR10 *DeltatxdOffsetlistpucchR10
}
