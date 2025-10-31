package ies

import "rrc/utils"

// N1PUCCH-AN-PersistentList ::= SEQUENCE OF utils.INTEGER // SIZE (1..4)
type N1pucchAnPersistentlist struct {
	Value []utils.INTEGER `lb:1,ub:4`
}
