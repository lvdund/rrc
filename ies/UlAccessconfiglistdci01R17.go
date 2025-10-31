package ies

import "rrc/utils"

// UL-AccessConfigListDCI-0-1-r17 ::= SEQUENCE OF utils.INTEGER // SIZE (1..3)
type UlAccessconfiglistdci01R17 struct {
	Value []utils.INTEGER `lb:1,ub:3`
}
