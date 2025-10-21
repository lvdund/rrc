package ies

import "rrc/utils"

// PLMN-IdentityList-NB-r13 ::= SEQUENCE OF PLMN-IdentityInfo-NB-r13
// SIZE (1..maxPLMN-r11)
type PlmnIdentitylistNbR13 struct {
	Value utils.Sequence[PlmnIdentityinfoNbR13]
}
