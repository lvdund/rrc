package ies

import "rrc/utils"

// PLMN-Identity-EUTRA-5GC ::= CHOICE
const (
	PlmnIdentityEutra5gcChoiceNothing = iota
	PlmnIdentityEutra5gcChoicePlmnIdentityEutra5gc
	PlmnIdentityEutra5gcChoicePlmnIndex
)

type PlmnIdentityEutra5gc struct {
	Choice               uint64
	PlmnIdentityEutra5gc *PlmnIdentity
	PlmnIndex            *utils.INTEGER `lb:1,ub:maxPLMN`
}
