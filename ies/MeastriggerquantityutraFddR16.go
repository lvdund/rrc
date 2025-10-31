package ies

import "rrc/utils"

// MeasTriggerQuantityUTRA-FDD-r16 ::= CHOICE
const (
	MeastriggerquantityutraFddR16ChoiceNothing = iota
	MeastriggerquantityutraFddR16ChoiceUtraFddRscpR16
	MeastriggerquantityutraFddR16ChoiceUtraFddEcn0R16
)

type MeastriggerquantityutraFddR16 struct {
	Choice         uint64
	UtraFddRscpR16 *utils.INTEGER `lb:-5,ub:91`
	UtraFddEcn0R16 *utils.INTEGER `lb:0,ub:49`
}
