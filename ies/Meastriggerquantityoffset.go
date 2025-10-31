package ies

import "rrc/utils"

// MeasTriggerQuantityOffset ::= CHOICE
const (
	MeastriggerquantityoffsetChoiceNothing = iota
	MeastriggerquantityoffsetChoiceRsrp
	MeastriggerquantityoffsetChoiceRsrq
	MeastriggerquantityoffsetChoiceSinr
)

type Meastriggerquantityoffset struct {
	Choice uint64
	Rsrp   *utils.INTEGER `lb:-30,ub:30`
	Rsrq   *utils.INTEGER `lb:-30,ub:30`
	Sinr   *utils.INTEGER `lb:-30,ub:30`
}
