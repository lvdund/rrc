package ies

import "rrc/utils"

// SRB-ToReleaseList-r15 ::= SEQUENCE OF utils.INTEGER // SIZE (1..2)
type SrbToreleaselistR15 struct {
	Value []utils.INTEGER `lb:1,ub:2`
}
