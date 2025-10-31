package ies

import "rrc/utils"

// UL-AccessConfigListDCI-0-1-r16 ::= SEQUENCE OF utils.INTEGER // SIZE (1..64)
type UlAccessconfiglistdci01R16 struct {
	Value []utils.INTEGER `lb:1,ub:64`
}
