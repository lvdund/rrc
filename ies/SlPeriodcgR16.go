package ies

import "rrc/utils"

// SL-PeriodCG-r16 ::= CHOICE
const (
	SlPeriodcgR16ChoiceNothing = iota
	SlPeriodcgR16ChoiceSlPeriodcg1R16
	SlPeriodcgR16ChoiceSlPeriodcg2R16
)

type SlPeriodcgR16 struct {
	Choice         uint64
	SlPeriodcg1R16 *utils.ENUMERATED
	SlPeriodcg2R16 *utils.INTEGER `lb:1,ub:99`
}
