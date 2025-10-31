package ies

import "rrc/utils"

// NextHopChainingCount ::= utils.INTEGER (0..7)
type Nexthopchainingcount struct {
	Value utils.INTEGER `lb:0,ub:7`
}
