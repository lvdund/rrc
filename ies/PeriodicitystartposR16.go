package ies

import "rrc/utils"

// PeriodicityStartPos-r16 ::= CHOICE
const (
	PeriodicitystartposR16ChoiceNothing = iota
	PeriodicitystartposR16ChoicePeriodicity10ms
	PeriodicitystartposR16ChoicePeriodicity20ms
	PeriodicitystartposR16ChoicePeriodicity40ms
	PeriodicitystartposR16ChoicePeriodicity80ms
	PeriodicitystartposR16ChoicePeriodicity160ms
	PeriodicitystartposR16ChoiceSpare3
	PeriodicitystartposR16ChoiceSpare2
	PeriodicitystartposR16ChoiceSpare1
)

type PeriodicitystartposR16 struct {
	Choice           uint64
	Periodicity10ms  *struct{}
	Periodicity20ms  *utils.INTEGER `lb:0,ub:1`
	Periodicity40ms  *utils.INTEGER `lb:0,ub:3`
	Periodicity80ms  *utils.INTEGER `lb:0,ub:7`
	Periodicity160ms *utils.INTEGER `lb:0,ub:15`
	Spare3           *struct{}
	Spare2           *struct{}
	Spare1           *struct{}
}
