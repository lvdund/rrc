package ies

import "rrc/utils"

// MRB-Identity-r17 ::= utils.INTEGER (1..512)
type MrbIdentityR17 struct {
	Value utils.INTEGER `lb:0,ub:512`
}
