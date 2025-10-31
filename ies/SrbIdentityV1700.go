package ies

import "rrc/utils"

// SRB-Identity-v1700 ::= utils.INTEGER (4)
type SrbIdentityV1700 struct {
	Value utils.INTEGER `lb:0,ub:4`
}
