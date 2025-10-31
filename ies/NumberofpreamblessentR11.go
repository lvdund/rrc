package ies

import "rrc/utils"

// NumberOfPreamblesSent-r11 ::= utils.INTEGER (1..200)
type NumberofpreamblessentR11 struct {
	Value utils.INTEGER `lb:0,ub:200`
}
