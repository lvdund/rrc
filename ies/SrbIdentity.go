package ies

import "rrc/utils"

// SRB-Identity ::= utils.INTEGER (1..3)
type SrbIdentity struct {
	Value utils.INTEGER `lb:0,ub:3`
}
