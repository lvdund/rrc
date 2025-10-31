package ies

import "rrc/utils"

// TMGI-r17-plmn-Id-r17 ::= CHOICE
const (
	TmgiR17PlmnIdR17ChoiceNothing = iota
	TmgiR17PlmnIdR17ChoicePlmnIndex
	TmgiR17PlmnIdR17ChoiceExplicitvalue
)

type TmgiR17PlmnIdR17 struct {
	Choice        uint64
	PlmnIndex     *utils.INTEGER `lb:1,ub:maxPLMN`
	Explicitvalue *PlmnIdentity
}
