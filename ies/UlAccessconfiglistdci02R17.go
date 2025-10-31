package ies

import "rrc/utils"

// UL-AccessConfigListDCI-0-2-r17 ::= SEQUENCE OF utils.INTEGER // SIZE (1..64)
type UlAccessconfiglistdci02R17 struct {
	Value []utils.INTEGER `lb:1,ub:64`
}
