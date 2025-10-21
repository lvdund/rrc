package ies

import "rrc/utils"

// UplinkPowerControlDedicatedSTTI-r15 ::= SEQUENCE
type UplinkpowercontroldedicatedsttiR15 struct {
	AccumulationenabledsttiR15  bool
	DeltatxdOffsetlistspucchR15 *DeltatxdOffsetlistspucchR15
	UplinkpowerCsipayload       bool
}
