package ies

import "rrc/utils"

// UplinkPowerControlDedicatedSTTI-r15 ::= SEQUENCE
type UplinkpowercontroldedicatedsttiR15 struct {
	AccumulationenabledsttiR15  utils.BOOLEAN
	DeltatxdOffsetlistspucchR15 *DeltatxdOffsetlistspucchR15
	UplinkpowerCsipayload       utils.BOOLEAN
}
