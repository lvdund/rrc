package ies

import "rrc/utils"

// N1PUCCH-AN-CS-r10 ::= SEQUENCE OF utils.INTEGER // SIZE (1..4)
type N1pucchAnCsR10 struct {
	Value []utils.INTEGER `lb:1,ub:4`
}
