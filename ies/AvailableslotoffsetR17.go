package ies

import "rrc/utils"

// AvailableSlotOffset-r17 ::= utils.INTEGER (0..7)
type AvailableslotoffsetR17 struct {
	Value utils.INTEGER `lb:0,ub:7`
}
