package ies

import "rrc/utils"

// TimeSinceFailure-r11 ::= utils.INTEGER (0..172800)
type TimesincefailureR11 struct {
	Value utils.INTEGER `lb:0,ub:172800`
}
