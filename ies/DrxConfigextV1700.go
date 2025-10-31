package ies

import "rrc/utils"

// DRX-ConfigExt-v1700 ::= SEQUENCE
type DrxConfigextV1700 struct {
	DrxHarqRttTimerdlR17 utils.INTEGER `lb:0,ub:448`
	DrxHarqRttTimerulR17 utils.INTEGER `lb:0,ub:448`
}
