package ies

import "rrc/utils"

// STAG-Id-r11 ::= utils.INTEGER (1..maxSTAG-r11)
type StagIdR11 struct {
	Value utils.INTEGER `lb:0,ub:maxSTAGR11`
}
