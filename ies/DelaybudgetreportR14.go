package ies

import "rrc/utils"

// DelayBudgetReport-r14 ::= CHOICE
const (
	DelaybudgetreportR14ChoiceNothing = iota
	DelaybudgetreportR14ChoiceType1
	DelaybudgetreportR14ChoiceType2
)

type DelaybudgetreportR14 struct {
	Choice uint64
	Type1  *utils.ENUMERATED
	Type2  *utils.ENUMERATED
}
