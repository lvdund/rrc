package ies

import "rrc/utils"

// DRB-Identity ::= utils.INTEGER (1..32)
type DrbIdentity struct {
	Value utils.INTEGER `lb:0,ub:32`
}
