package ies

import "rrc/utils"

// TimeSinceFailure-r16 ::= utils.INTEGER (0..172800)
type TimesincefailureR16 struct {
	Value utils.INTEGER `lb:0,ub:172800`
}
