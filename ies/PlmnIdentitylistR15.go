package ies

import "rrc/utils"

// PLMN-IdentityList-r15 ::= SEQUENCE OF PLMN-IdentityInfo-r15
// SIZE (1..maxPLMN-r11)
type PlmnIdentitylistR15 struct {
	Value utils.Sequence[PlmnIdentityinfoR15]
}
