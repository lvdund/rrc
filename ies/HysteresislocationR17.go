package ies

import "rrc/utils"

// HysteresisLocation-r17 ::= utils.INTEGER (0..32768)
type HysteresislocationR17 struct {
	Value utils.INTEGER `lb:0,ub:32768`
}
