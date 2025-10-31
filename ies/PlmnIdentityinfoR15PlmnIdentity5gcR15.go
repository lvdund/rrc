package ies

import "rrc/utils"

// PLMN-IdentityInfo-r15-plmn-Identity-5GC-r15 ::= CHOICE
const (
	PlmnIdentityinfoR15PlmnIdentity5gcR15ChoiceNothing = iota
	PlmnIdentityinfoR15PlmnIdentity5gcR15ChoicePlmnIdentityR15
	PlmnIdentityinfoR15PlmnIdentity5gcR15ChoicePlmnIndexR15
)

type PlmnIdentityinfoR15PlmnIdentity5gcR15 struct {
	Choice          uint64
	PlmnIdentityR15 *PlmnIdentity
	PlmnIndexR15    *utils.INTEGER `lb:1,ub:maxPLMNR11`
}
