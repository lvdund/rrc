package ies

import "rrc/utils"

// UL-AccessConfigListDCI-1-1-r16 ::= SEQUENCE OF utils.INTEGER // SIZE (1..16)
type UlAccessconfiglistdci11R16 struct {
	Value []utils.INTEGER `lb:1,ub:16`
}
