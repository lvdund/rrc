package ies

import "rrc/utils"

// SL-ResourceReservePeriod-r16 ::= CHOICE
const (
	SlResourcereserveperiodR16ChoiceNothing = iota
	SlResourcereserveperiodR16ChoiceSlResourcereserveperiod1R16
	SlResourcereserveperiodR16ChoiceSlResourcereserveperiod2R16
)

type SlResourcereserveperiodR16 struct {
	Choice                      uint64
	SlResourcereserveperiod1R16 *utils.ENUMERATED
	SlResourcereserveperiod2R16 *utils.INTEGER `lb:1,ub:99`
}
