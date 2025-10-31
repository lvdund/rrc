package ies

import "rrc/utils"

// SR-NPRACH-Resource-NB-r15-nprach-SubCarrierIndex-r15 ::= CHOICE
const (
	SrNprachResourceNbR15NprachSubcarrierindexR15ChoiceNothing = iota
	SrNprachResourceNbR15NprachSubcarrierindexR15ChoiceNprachFmt0fmt1R15
	SrNprachResourceNbR15NprachSubcarrierindexR15ChoiceNprachFmt2R15
)

type SrNprachResourceNbR15NprachSubcarrierindexR15 struct {
	Choice            uint64
	NprachFmt0fmt1R15 *utils.INTEGER `lb:0,ub:47`
	NprachFmt2R15     *utils.INTEGER `lb:0,ub:143`
}
