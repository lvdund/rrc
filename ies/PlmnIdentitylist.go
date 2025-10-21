package ies

import "rrc/utils"

// PLMN-IdentityList ::= SEQUENCE OF PLMN-IdentityInfo
// SIZE (1..maxPLMN-r11)
type PlmnIdentitylist struct {
	Value utils.Sequence[PlmnIdentityinfo]
}
