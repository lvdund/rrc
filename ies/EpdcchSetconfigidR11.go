package ies

import "rrc/utils"

// EPDCCH-SetConfigId-r11 ::= utils.INTEGER (0..1)
type EpdcchSetconfigidR11 struct {
	Value utils.INTEGER `lb:0,ub:1`
}
