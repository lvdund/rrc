package ies

import "rrc/utils"

// PMCH-Config-r12-dataMCS-r12 ::= CHOICE
const (
	PmchConfigR12DatamcsR12ChoiceNothing = iota
	PmchConfigR12DatamcsR12ChoiceNormalR12
	PmchConfigR12DatamcsR12ChoiceHigerorderR12
)

type PmchConfigR12DatamcsR12 struct {
	Choice        uint64
	NormalR12     *utils.INTEGER `lb:0,ub:28`
	HigerorderR12 *utils.INTEGER `lb:0,ub:27`
}
