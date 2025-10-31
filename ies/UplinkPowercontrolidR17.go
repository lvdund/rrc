package ies

import "rrc/utils"

// Uplink-powerControlId-r17 ::= utils.INTEGER (1.. maxUL-TCI-r17)
type UplinkPowercontrolidR17 struct {
	Value utils.INTEGER `lb:0,ub:maxULTciR17`
}
