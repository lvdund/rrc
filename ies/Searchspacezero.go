package ies

import "rrc/utils"

// SearchSpaceZero ::= utils.INTEGER (0..15)
type Searchspacezero struct {
	Value utils.INTEGER `lb:0,ub:15`
}
