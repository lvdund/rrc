package ies

import "rrc/utils"

// DRX-Config-v1310-longDRX-CycleStartOffset-v1310 ::= SEQUENCE
type DrxConfigV1310LongdrxCyclestartoffsetV1310 struct {
	Sf60V1310 utils.INTEGER `lb:0,ub:59`
}
