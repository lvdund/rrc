package ies

import "rrc/utils"

// PLMN-IdentityList-MBMS-r14 ::= SEQUENCE OF PLMN-Identity
// SIZE (1..maxPLMN-r11)
type PlmnIdentitylistMbmsR14 struct {
	Value utils.Sequence[PlmnIdentity]
}
