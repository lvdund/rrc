package ies

import "rrc/utils"

// PLMN-IdentityInfo2-r12 ::= CHOICE
const (
	PlmnIdentityinfo2R12ChoiceNothing = iota
	PlmnIdentityinfo2R12ChoicePlmnIndexR12
	PlmnIdentityinfo2R12ChoicePlmnidentityR12
)

type PlmnIdentityinfo2R12 struct {
	Choice          uint64
	PlmnIndexR12    *utils.INTEGER `lb:1,ub:maxPLMNR11`
	PlmnidentityR12 *PlmnIdentity
}
