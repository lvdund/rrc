package ies

import "rrc/utils"

// MBMS-SAI-r11 ::= utils.INTEGER (0..65535)
type MbmsSaiR11 struct {
	Value utils.INTEGER `lb:0,ub:65535`
}
