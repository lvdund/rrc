package ies

import "rrc/utils"

// PLMN-IdentityInfo-5GC-NB-r16-plmn-Identity-5GC-r16 ::= CHOICE
const (
	PlmnIdentityinfo5gcNbR16PlmnIdentity5gcR16ChoiceNothing = iota
	PlmnIdentityinfo5gcNbR16PlmnIdentity5gcR16ChoicePlmnIdentityR16
	PlmnIdentityinfo5gcNbR16PlmnIdentity5gcR16ChoicePlmnIndexR16
)

type PlmnIdentityinfo5gcNbR16PlmnIdentity5gcR16 struct {
	Choice          uint64
	PlmnIdentityR16 *PlmnIdentity
	PlmnIndexR16    *utils.INTEGER `lb:1,ub:maxPLMNR11`
}
