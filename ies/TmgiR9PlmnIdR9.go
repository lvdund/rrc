package ies

import "rrc/utils"

// TMGI-r9-plmn-Id-r9 ::= CHOICE
const (
	TmgiR9PlmnIdR9ChoiceNothing = iota
	TmgiR9PlmnIdR9ChoicePlmnIndexR9
	TmgiR9PlmnIdR9ChoiceExplicitvalueR9
)

type TmgiR9PlmnIdR9 struct {
	Choice          uint64
	PlmnIndexR9     *utils.INTEGER `lb:1,ub:maxPLMNR11`
	ExplicitvalueR9 *PlmnIdentity
}
