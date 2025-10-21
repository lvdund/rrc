package ies

import "rrc/utils"

// PLMN-IdentityList4-r12 ::= SEQUENCE OF PLMN-IdentityInfo2-r12
// SIZE (1..maxPLMN-r11)
type PlmnIdentitylist4R12 struct {
	Value utils.Sequence[PlmnIdentityinfo2R12]
}
