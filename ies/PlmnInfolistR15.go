package ies

import "rrc/utils"

// PLMN-InfoList-r15 ::= SEQUENCE OF PLMN-Info-r15
// SIZE (1..maxPLMN-r11)
type PlmnInfolistR15 struct {
	Value utils.Sequence[PlmnInfoR15]
}
