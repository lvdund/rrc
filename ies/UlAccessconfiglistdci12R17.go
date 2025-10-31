package ies

import "rrc/utils"

// UL-AccessConfigListDCI-1-2-r17 ::= SEQUENCE OF utils.INTEGER // SIZE (1..16)
type UlAccessconfiglistdci12R17 struct {
	Value []utils.INTEGER `lb:1,ub:16`
}
