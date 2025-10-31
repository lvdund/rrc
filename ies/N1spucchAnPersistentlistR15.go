package ies

import "rrc/utils"

// N1SPUCCH-AN-PersistentList-r15 ::= SEQUENCE OF utils.INTEGER // SIZE (1..4)
type N1spucchAnPersistentlistR15 struct {
	Value []utils.INTEGER `lb:1,ub:4`
}
